package v1

import "github.com/kubescape/go-git-url/apis/azureapi"

type AzureURL struct {
	host     string
	owner    string // repo owner
	repo     string // repo name
	project  string
	branch   string
	tag      string
	path     string
	token    string // azure token
	isFile   bool   // is the URL is pointing to a file or not
	// username string
	// password string

	azureAPI azureapi.IAzureAPI
}
