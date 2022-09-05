package v1

import "github.com/kubescape/go-git-url/apis/githubapi"

type GitHubURL struct {
	host   string // default is github
	owner  string // repo owner
	repo   string // repo name
	branch string
	path   string
	token  string // github token
	isFile bool   // is the URL is pointing to a file or not

	gitHubAPI githubapi.IGitHubAPI
}
