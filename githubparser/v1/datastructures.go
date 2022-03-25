package v1

type GitHubURL struct {
	host   string // default is github
	owner  string // repo owner
	repo   string // repo name
	branch string
	path   string
	token  string // github token
	isFile bool   // is the URL is pointing to a file or not
}
