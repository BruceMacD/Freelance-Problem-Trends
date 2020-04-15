package analyzer

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
)

// Analyzer that reads data from a database into data structures
type Analyzer struct {
	DB      *sql.DB
	Verbose bool
}

// GetFilteredWordsByOccurance returns a WordList of occurances, but removes common words that arent interesting (such as "the")
func (a *Analyzer) GetFilteredWordsByOccurance(f map[string]bool) WordList {
	return a.getWordsByOccurance(f)
}

// GetWordsByOccurance returns words in database sorted by occurance
func (a *Analyzer) getWordsByOccurance(filter map[string]bool) WordList {
	wos := a.getAllWords()
	wbo := mapWordsByOccurance(wos)
	if a.Verbose {
		fmt.Printf("Number of unique words from all postings: %d\n", len(wbo))
	}

	// remove filtered words from map
	for rem := range filter {
		delete(wbo, rem)
	}

	// convert the words by occurance map to a sorted list of Words by occurances
	wl := make(WordList, len(wbo))
	i := 0
	for val, occ := range wbo {
		wl[i] = Word{
			Value:      val,
			Occurances: occ,
		}
		i++
	}
	sort.Sort(sort.Reverse(wl))

	return wl
}

// converts the database into a slice of individual words
func (a *Analyzer) getAllWords() (wos []string) {
	ts := a.getAllTitlesFromDB()
	tsw := parseAllContentsToSingleWords(ts)
	if a.Verbose {
		fmt.Printf("Number of words from posting titles: %d\n", len(tsw))
	}

	ds := a.getAllDescriptionsFromDB()
	dsw := parseAllContentsToSingleWords(ds)
	if a.Verbose {
		fmt.Printf("Number of words from posting descriptions: %d\n", len(dsw))
	}

	wos = append(wos, tsw...)
	wos = append(wos, dsw...)

	return
}

func (a *Analyzer) getAllTitlesFromDB() (ts []string) {
	var title string
	rows, _ := a.DB.Query("SELECT title FROM postings")
	for rows.Next() {
		rows.Scan(&title)
		ts = append(ts, title)
	}
	return
}

func (a *Analyzer) GetAllSkillsFromDB() (s []string) {
	var ids []string
	var id string

	// read IDs into memory
	rows, _ := a.DB.Query("SELECT id FROM postings")
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids, id)
	}

	// parse skill from ID
	for _, path := range ids {
		split := strings.Split(path, "/")
		if split != nil && split[0] == "projects" {
			ids = append(ids, split[1])
		}
	}

	fmt.Println(ids)

	return
}

func (a *Analyzer) getAllDescriptionsFromDB() (ds []string) {
	var desc string

	rows, _ := a.DB.Query("SELECT description FROM postings")
	for rows.Next() {
		rows.Scan(&desc)
		ds = append(ds, desc)
	}
	return
}

// converts a slice of multi-word sentences to single words
func parseAllContentsToSingleWords(mw []string) (sw []string) {
	for _, s := range mw {
		words := strings.Split(s, " ")
		for _, w := range words {
			sw = append(sw, w)
		}
	}
	return
}

func mapWordsByOccurance(ws []string) (wcs map[string]int) {
	wcs = make(map[string]int)
	for _, w := range ws {
		lw := strings.ToLower(w)
		numOcc := wcs[lw] + 1
		wcs[lw] = numOcc
	}
	return
}
