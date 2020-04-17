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

		// what are employers looking for in their job candidates?
		analyzeTopRequirements(analyzer)
		analyzeTopTenSkills(analyzer)
		analyzeAllTopSkills(analyzer)
		// TODO: heat map skills to requirements
	}

	fmt.Printf("Program completed successfully.\n")
}

func analyzeTopRequirements(a analyzer.Analyzer) {
	wl := a.GetFilteredWordsByOccurance(analyzer.GetCommonWords())

	if a.Verbose && len(wl) > 1000 {
		for i := 0; i < 1000; i++ {
			fmt.Printf("\"%s\": %d,\n", wl[i].Value, wl[i].Occurances)
		}
	}

	visualizer.DrawTopNOccuranceBarChart(wl, "Top Description Job Requirements", "top_requirements_bar", 10)
}

func analyzeTopTenSkills(a analyzer.Analyzer) {
	sl := a.GetSortedSkillsByOccurance()

	if a.Verbose && len(sl) > 10 {
		for i := 0; i < 10; i++ {
			fmt.Printf("\"%s\": %d,\n", sl[i].Value, sl[i].Occurances)
		}
	}

	visualizer.DrawTopNOccuranceBarChart(sl, "Top Job Skills", "top_skills_bar", 10)
}

func analyzeAllTopSkills(a analyzer.Analyzer) {
	sl := a.GetSkillsByOccurance()
	if a.Verbose && len(sl) > 10 {
		fmt.Printf("Total num skills: %d\n", len(sl))
	}
	visualizer.DrawWordCloud(sl, "All Job Skills", "all_skills_cloud")
}
