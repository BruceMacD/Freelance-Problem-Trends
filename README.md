# Freelance-Problem-Trends
Crawling and parsing postings on freelance job boards to find trends.

This project was used to collect and visualize the data used in the accompanying blog post:

https://medium.com/@brucewmacdonald/what-are-the-most-in-demand-skills-for-remote-freelance-work-bafff40c07f7

# Running
```
> go build
> ./Freelance-Problem-Trends --help
Usage of ./Freelance-Problem-Trends:
  -analyze
        indicates if the data should be read into memory from the database
  -crawl
        indicates if results should be crawled from freelancer site
  -endPage int
        the page to stop crawling at
  -startPage int
        the page to start crawling from
```
