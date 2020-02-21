package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/BruceMacD/Freelance-Problem-Trends/crawler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	doCrawl := flag.Bool("crawl", false, "indicates if results should be crawled from freelancer site")
	startPage := flag.Int("startPage", 0, "the page to start crawling from")
	endPage := flag.Int("endPage", 0, "the page to stop crawling at")
	flag.Parse()

	database, err := sql.Open("sqlite3", "./freelancer.db")
	if err != nil {
		fmt.Printf("Could not open sqlite database: %s\n", err)
		os.Exit(1)
	}

	if *doCrawl {
		crawler := crawler.FreelanceCrawler{DB: database}
		crawler.CrawlAndWrite(*startPage, *endPage)
	}

	fmt.Printf("Program completed successfully.\n")
}
