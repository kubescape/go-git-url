package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/armosec/url-git-go/apis"
)

const (
	DEFAULT_HOST string = "github.com"
	RAW_HOST     string = "raw.githubusercontent.com"
)

type IGitHubAPI interface {
	GetRepoTree(owner, repo, branch string, headres *Headres) (*Tree, error)
	GetDefaultBranchName(owner, repo string, headres *Headres) (string, error)
}
type GitHubAPI struct {
	httpClient *http.Client
}

func NewGitHubAPI() *GitHubAPI { return &GitHubAPI{httpClient: &http.Client{}} }

func (gh *GitHubAPI) GetRepoTree(owner, repo, branch string, headres *Headres) (*Tree, error) {
	treeAPI := APIRepoTree(owner, repo, branch)
	body, err := apis.HttpGet(gh.httpClient, treeAPI, headres.ToMap())
	if err != nil {
		return nil, err
	}

	var tree Tree
	err = json.Unmarshal([]byte(body), &tree)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", treeAPI, err.Error())
	}
	return &tree, nil

}

func (gh *GitHubAPI) GetDefaultBranchName(owner, repo string, headres *Headres) (string, error) {

	body, err := apis.HttpGet(gh.httpClient, APIDefaultBranch(owner, repo), headres.ToMap())
	if err != nil {
		return "", err
	}

	var data githubDefaultBranchAPI
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}
	return data.DefaultBranch, nil
}

func APIRepoTree(owner, repo, branch string) string {
	return fmt.Sprintf("https://api.%s/repos/%s/%s/git/trees/%s?recursive=1", DEFAULT_HOST, owner, repo, branch)
}

func APIRaw(owner, repo, branch, path string) string {
	return fmt.Sprintf("https://%s/%s/%s/%s/%s", RAW_HOST, owner, repo, branch, path)
}
func APIDefaultBranch(owner, repo string) string {
	return fmt.Sprintf("https://api.%s/repos/%s/%s", DEFAULT_HOST, owner, repo)
}
