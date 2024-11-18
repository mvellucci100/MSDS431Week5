package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
)

func TestScraper(t *testing.T) {
	// Mock data for URLs and their responses
	mockResponses := map[string]string{
		"/page1": "This is the content of page 1.",
		"/page2": "This is the content of page 2.",
	}

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response, exists := mockResponses[r.URL.Path]; exists {
			fmt.Fprintf(w, "<body>%s</body>", response)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Define the test URLs based on the mock server
	urls := []string{
		server.URL + "/page1",
		server.URL + "/page2",
	}

	// Create a Colly collector
	c := colly.NewCollector()

	// Slice to store the results
	var results []map[string]string

	// Handle the "body" HTML element
	c.OnHTML("body", func(e *colly.HTMLElement) {
		results = append(results, map[string]string{
			"url":     e.Request.URL.String(),
			"content": e.Text,
		})
	})

	// Iterate over the URLs and visit them
	for _, url := range urls {
		err := c.Visit(url)
		assert.NoError(t, err, "Visiting URL should not return an error")
	}

	// Expected results
	expected := []map[string]string{
		{"url": urls[0], "content": "This is the content of page 1."},
		{"url": urls[1], "content": "This is the content of page 2."},
	}

	// Assert that the results match the expected data
	assert.Equal(t, expected, results, "The scraped results should match the expected output")
}
