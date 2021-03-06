package analyzer

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

// GetSortedSkillsByOccurance parses skills from freelance postings and returns a sorted WordList by occurance
func (a *Analyzer) GetSortedSkillsByOccurance() WordList {
	so := a.GetSkillsByOccurance()
	if a.Verbose {
		fmt.Printf("Number of unique skills from all postings: %d\n", len(so))
	}

	return getSortedWordList(so)
}

// GetSkillsByOccurance parses skills from freelance postings and returns a map of occurances
func (a *Analyzer) GetSkillsByOccurance() map[string]int {
	s := a.getAllSkillsFromDB()
	return mapWordsByOccurance(s)
}

// GetSkillsByRequirements matches which top reqiurements are related to which top skills
func (a *Analyzer) GetSkillsByRequirements() (skills, requirements []string, hmd [][3]interface{}, max int, e error) {
	s := a.GetSortedSkillsByOccurance()
	r := a.GetFilteredWordsByOccurance(GetCommonWords())

	if len(s) < 10 || len(r) < 10 {
		e = errors.New("There isn't enough data to analyze skill requirements")
		return nil, nil, nil, 0, e
	}

	// get 10 skills from s
	// get 10 requirements from r
	for i := 0; i < 10; i++ {
		skills = append(skills, s[i].Value)
		requirements = append(requirements, r[i].Value)
	}
	// for each skill description count how many occurance of requirement
	// this could be faster, but meh
	// {skill_index, req_index, value}
	// first skill/req = 0-9
	// second skill/req = 10-11
	// ..etc
	hmd = make([][3]interface{}, 0)
	for i, skill := range skills {
		desc := a.getDescriptionsForSkillFromDB(skill)
		for j, req := range requirements {
			cnt := 0
			for _, d := range desc {
				if strings.Contains(d, req) {
					cnt++
				}
			}
			hmd = append(hmd, [3]interface{}{i, j, cnt})
			if cnt > max {
				max = cnt
			}
		}
	}

	return skills, requirements, hmd, max, nil
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
	rows, err := a.DB.Query("SELECT title FROM postings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
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
	rows, err := a.DB.Query("SELECT id FROM postings")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
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

	rows, err := a.DB.Query("SELECT description FROM postings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&desc)
		ds = append(ds, desc)
	}
	return
}

func (a *Analyzer) getDescriptionsForSkillFromDB(skill string) (ds []string) {
	var desc string

	rows, err := a.DB.Query("SELECT description FROM postings WHERE id LIKE ?", "%"+skill+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&desc)
		ds = append(ds, desc)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
