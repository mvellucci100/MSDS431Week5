package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"

    "github.com/gocolly/colly/v2"
)

func main() {
    // Wikipedia URLs for topic of interest
    urls := []string{
        "https://en.wikipedia.org/wiki/Robotics",
        "https://en.wikipedia.org/wiki/Robot",
        "https://en.wikipedia.org/wiki/Reinforcement_learning",
        "https://en.wikipedia.org/wiki/Robot_Operating_System",
        "https://en.wikipedia.org/wiki/Intelligent_agent",
        "https://en.wikipedia.org/wiki/Software_agent",
        "https://en.wikipedia.org/wiki/Robotic_process_automation",
        "https://en.wikipedia.org/wiki/Chatbot",
        "https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
        "https://en.wikipedia.org/wiki/Android_(robot)",
    }

    // Create a new collector
    c := colly.NewCollector()

    // Create a slice to hold the scraped data
    var results []map[string]string

    // Define what to do when a visited HTML element is found
    c.OnHTML("body", func(e *colly.HTMLElement) {
        text := strings.TrimSpace(e.Text)
        results = append(results, map[string]string{"url": e.Request.URL.String(), "content": text})
    })

    // Iterate over the URLs and visit them
    for _, url := range urls {
        err := c.Visit(url)
        if err != nil {
            fmt.Printf("Failed to visit %s: %v\n", url, err)
        }
    }

    // Create and open the JSON lines file
    file, err := os.Create("scraped_data.jsonl")
    if err != nil {
        fmt.Printf("Could not create file: %v\n", err)
        return
    }
    defer file.Close()

    // Write results to the JSON lines file
    for _, result := range results {
        line, err := json.Marshal(result)
        if err != nil {
            fmt.Printf("Could not marshal JSON: %v\n", err)
            continue
        }
        file.WriteString(string(line) + "\n")
    }

    fmt.Println("Scraping completed. Data saved to scraped_data.jsonl.")
}
