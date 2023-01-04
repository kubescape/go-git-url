package bitbucketapi

import "time"

type ObjectType string

type Tree struct{}

type bitbucketBranchingModel struct {
	Development struct {
		Name string `json:"name"`
	} `json:"development"`
}

type Headers struct {
	Token string
}

type Commits struct {
	Values  []Commit `json:"values"`
	Pagelen int      `json:"pagelen"`
	Next    string   `json:"next"`
}

type Commit struct {
	Type   string    `json:"type"`
	Hash   string    `json:"hash"`
	Date   time.Time `json:"date"`
	Author struct {
		Type string `json:"type"`
		Raw  string `json:"raw"`
	} `json:"author"`
	Message string `json:"message"`
}
