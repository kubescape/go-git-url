package v1

type GitLabURL struct {
	host    string // default is github
	owner   string // repo owner
	repo    string // repo name
	project string
	branch  string
	path    string
	token   string // github token
}
