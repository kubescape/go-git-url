package v1

import (
	"github.com/kubescape/go-git-url/apis/bitbucketapi"
)

type BitBucketURL struct {
	host    string
	owner   string // repo owner
	repo    string // repo name
	project string
	branch  string
	path    string
	token   string // github token
	isFile  bool

	bitBucketAPI bitbucketapi.IBitBucketAPI
}
