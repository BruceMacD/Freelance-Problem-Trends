package crawler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
)

const freelancerURL string = "https://www.freelancer.com/jobs/"
const freelanceItemID string = ".JobSearchCard-item "
const freelanceTitleID string = ".JobSearchCard-primary-heading-link"
const freelanceDescriptionID string = ".JobSearchCard-primary-description"
const freelanceLinkID string = "JobSearchCard-ctas-btn btn btn-mini btn-success"

// FreelanceCrawler that writes scraped information to specified file
type FreelanceCrawler struct {
	DB    *sql.DB
	Debug bool
}

func (crawler *FreelanceCrawler) savePost(post Posting) {
	statement, _ := crawler.DB.Prepare("INSERT INTO postings (id, title, description) VALUES (?, ?, ?)")
	statement.Exec(post.link, post.title, post.description)
}

func (crawler *FreelanceCrawler) start(startPage int, endPage int) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(freelanceItemID, func(e *colly.HTMLElement) {
		post := Posting{}
		post.title = e.ChildText(freelanceTitleID)
		post.description = e.ChildText(freelanceDescriptionID)
		// all postings have a unique link
		post.link = e.ChildAttr("a", "href")

		if crawler.Debug {
			fmt.Println(post)
		}

		crawler.savePost(post)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Limit the number of threads started by colly to two
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       2 * time.Second,
	})

	for i := startPage; i <= endPage; i++ {
		c.Visit(fmt.Sprintf("%s%d", freelancerURL, i))
	}

	c.Wait()
}

// CrawlAndWrite crawls the Freelancer job postings from the given page and write the contents to the file
func (crawler *FreelanceCrawler) CrawlAndWrite(startPage int, endPage int) {
	if startPage == 0 || endPage == 0 {
		fmt.Printf("Start and/or end pages not set. Not crawling.\n")
		return
	}
	statement, _ := crawler.DB.Prepare("CREATE TABLE IF NOT EXISTS postings (id TEXT PRIMARY KEY, title TEXT, description TEXT)")
	statement.Exec()
	crawler.start(startPage, endPage)
}
