package analyzer

// Word contains the value of the word and the number of times is was encountered during processing
type Word struct {
	Value      string
	Occurances int
}

// WordList is a list of Words
type WordList []Word

func (w WordList) Len() int           { return len(w) }
func (w WordList) Less(i, j int) bool { return w[i].Occurances < w[j].Occurances }
func (w WordList) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
