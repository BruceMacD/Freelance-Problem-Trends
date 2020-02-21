package crawler

import "fmt"

// Posting stores the basic data assosiated with a job post
type Posting struct {
	link        string
	title       string
	description string
}

func (post Posting) String() string {
	return fmt.Sprintf("[%s, %s, %s]", post.link, post.title, post.description)
}
