package bitbucketapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kubescape/go-git-url/apis"
)

const (
	DEFAULT_HOST string = "bitbucket.org"
)

type IBitBucketAPI interface {
	GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error)
	GetDefaultBranchName(owner, repo string, headers *Headers) (string, error)
	GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error)
}

type BitBucketAPI struct {
	httpClient *http.Client
}

func NewBitBucketAPI() *BitBucketAPI { return &BitBucketAPI{httpClient: &http.Client{}} }

func (gl *BitBucketAPI) GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error) {
	//TODO implement me
	return nil, fmt.Errorf("GetRepoTree is not supported")
}

func (gl *BitBucketAPI) GetDefaultBranchName(owner, repo string, headers *Headers) (string, error) {
	body, err := apis.HttpGet(gl.httpClient, APIBranchingModel(owner, repo), headers.ToMap())
	if err != nil {
		return "", err
	}

	var data bitbucketBranchingModel
	if err = json.Unmarshal([]byte(body), &data); err != nil {
		return "", err
	}

	return data.Development.Name, nil
}

func (gl *BitBucketAPI) GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error) {
	body, err := apis.HttpGet(gl.httpClient, APILastCommitsOfBranch(owner, repo, branch), headers.ToMap())
	if err != nil {
		return nil, err
	}

	var data Commits
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}
	return &data.Values[0], nil
}

// APIBranchingModel
// API Ref: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-branching-model/#api-group-branching-model
// Example: https://bitbucket.org/!api/2.0/repositories/matthyx/ks-testing-public/branching-model
func APIBranchingModel(owner, repo string) string {
	p, _ := url.JoinPath("!api/2.0/repositories", owner, repo, "branching-model")
	u := url.URL{
		Scheme: "https",
		Host:   DEFAULT_HOST,
		Path:   p,
	}
	return u.String()
}

// APILastCommitsOfBranch
// API Ref: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-commits-get
// Example: https://bitbucket.org/!api/2.0/repositories/matthyx/ks-testing-public/commits/?include=master
func APILastCommitsOfBranch(owner, repo, branch string) string {
	p, _ := url.JoinPath("!api/2.0/repositories", owner, repo, "commits")
	u := url.URL{
		Scheme:   "https",
		Host:     DEFAULT_HOST,
		Path:     p,
		RawQuery: fmt.Sprintf("include=%s", branch),
	}
	return u.String()
}
