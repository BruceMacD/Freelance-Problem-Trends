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
	wos := a.getAllWords()
	wbo := mapWordsByOccurance(wos)
	if a.Verbose {
		fmt.Printf("Number of unique words from all postings: %d\n", len(wbo))
	}

	// remove filtered words from map
	for rem := range f {
		delete(wbo, rem)
	}

	return getSortedWordList(wbo)
}

// GetSkillsByOccurance parses skills from freelance postings and returns a sorted WordList by occurance
func (a *Analyzer) GetSkillsByOccurance() WordList {
	s := a.getAllSkillsFromDB()
	so := mapWordsByOccurance(s)
	if a.Verbose {
		fmt.Printf("Number of unique skills from all postings: %d\n", len(so))
	}

	return getSortedWordList(so)
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

func (a *Analyzer) getAllSkillsFromDB() (s []string) {
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

		if split != nil && len(split) > 3 && split[1] == "projects" {
			s = append(s, split[2])
		}
	}

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

func getSortedWordList(wcs map[string]int) (wl WordList) {
	wl = make(WordList, len(wcs))
	i := 0
	for val, occ := range wcs {
		wl[i] = Word{
			Value:      val,
			Occurances: occ,
		}
		i++
	}
	sort.Sort(sort.Reverse(wl))
	return
}
