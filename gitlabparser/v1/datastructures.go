package v1

import "github.com/kubescape/go-git-url/apis/gitlabapi"

type GitLabURL struct {
	host    string
	owner   string // repo owner
	repo    string // repo name
	project string
	branch  string
	path    string
	token   string // github token
	isFile  bool

	gitLabAPI gitlabapi.IGitLabAPI
}
