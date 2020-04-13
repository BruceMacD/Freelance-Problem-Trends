package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/BruceMacD/Freelance-Problem-Trends/analyzer"
	"github.com/BruceMacD/Freelance-Problem-Trends/crawler"
	"github.com/BruceMacD/Freelance-Problem-Trends/visualizer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	crawl := flag.Bool("crawl", false, "indicates if results should be crawled from freelancer site")
	startPage := flag.Int("startPage", 0, "the page to start crawling from")
	endPage := flag.Int("endPage", 0, "the page to stop crawling at")
	analyze := flag.Bool("analyze", false, "indicates if the data should be read into memory from the database and visualized")
	verbose := flag.Bool("v", false, "verbose output for debugging")
	flag.Parse()

	database, err := sql.Open("sqlite3", "./freelancer.db")
	if err != nil {
		fmt.Printf("Could not open sqlite database: %s\n", err)
		os.Exit(1)
	}

	if *crawl {
		crawler := crawler.FreelanceCrawler{DB: database, Verbose: *verbose}
		crawler.CrawlAndWrite(*startPage, *endPage)
	}

	if *analyze {
		analyzer := analyzer.Analyzer{DB: database, Verbose: *verbose}
		wl := analyzer.GetFilteredWordsByOccurance()

		if *verbose {
			for i := 0; i < 1000; i++ {
				fmt.Printf("\"%s\": %d,\n", wl[i].Value, wl[i].Occurances)
			}
		}

		visualizer.DrawTopNOccuranceBarChart(wl, "Keywords by Occurance", "top_ten_occurance_bar", 10)
	}

	fmt.Printf("Program completed successfully.\n")
}
