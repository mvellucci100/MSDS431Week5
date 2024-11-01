# MSDS431 
#Week 5 Assignment: Crawling and Scraping the Web

# Go Web Crawler/Scraper

This project is a web crawler and scraper developed in Go using the Colly framework. It retrieves text information from specified Wikipedia pages related to intelligent systems and robotics, processes the data, and outputs it in JSON lines format.

## Features

- Concurrent web crawling and scraping using Go's goroutines.
- Extracts and cleans text content from Wikipedia pages.
- Saves the extracted data in a JSON lines file for easy processing.

## Getting Started

### Prerequisites

- Go (1.16 or later) installed on your machine.
- Basic knowledge of Go programming and command line usage.

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/mvellucci100/MSDS431Week5.git
   cd MSDS431Week5
2. Initialize the Go Module:
   go mod init go-web-scraper
4. Install Colly framework:
   go get -u github.com/gocolly/colly/v2
5. Run the application
   go run main.go


