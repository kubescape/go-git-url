package apis

import "time"

type Commit struct {
	SHA       string    `json:"sha"`
	Author    Committer `json:"author"`
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
	Files     []Files   `json:"files"`
}

type Files struct {
	FileSHA  string `json:"sha"`
	Filename string `json:"filename"`
}

type Committer struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}
