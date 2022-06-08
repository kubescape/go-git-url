package v1

type AzureURL struct {
	host     string // default is github
	owner    string // repo owner
	repo     string // repo name
	project  string
	branch   string
	tag      string
	path     string
	token    string // github token
	isFile   bool   // is the URL is pointing to a file or not
	username string
	password string
}
