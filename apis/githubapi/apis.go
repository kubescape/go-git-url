package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kubescape/go-git-url/apis"
)

const (
	DEFAULT_HOST string = "github.com"
	RAW_HOST     string = "raw.githubusercontent.com"
)

type IGitHubAPI interface {
	GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error)
	GetDefaultBranchName(owner, repo string, headers *Headers) (string, error)
	GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error)
	// GetCommits(owner, repo, branch string, headers *Headers) ([]Commit, error)
	// GetLatestPathCommit(owner, repo, branch, fullPath string, headers *Headers) ([]]Commit, error)
}
type GitHubAPI struct {
	httpClient *http.Client
}

func NewGitHubAPI() *GitHubAPI { return &GitHubAPI{httpClient: &http.Client{}} }

func (gh *GitHubAPI) GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error) {
	treeAPI := APIRepoTree(owner, repo, branch)
	body, err := apis.HttpGet(gh.httpClient, treeAPI, headers.ToMap())
	if err != nil {
		return nil, err
	}

	var tree Tree
	if err = json.Unmarshal([]byte(body), &tree); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", treeAPI, err.Error())
	}
	return &tree, nil

}

func (gh *GitHubAPI) GetDefaultBranchName(owner, repo string, headers *Headers) (string, error) {

	body, err := apis.HttpGet(gh.httpClient, APIMetadata(owner, repo), headers.ToMap())
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

func (gh *GitHubAPI) GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error) {

	body, err := apis.HttpGet(gh.httpClient, APILastCommitsOfBranch(owner, repo, branch), headers.ToMap())
	if err != nil {
		return nil, err
	}

	var data Commit
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return &data, err
	}
	return &data, nil
}

// Get latest commit data of path/file
func (gh *GitHubAPI) GetFileLatestCommit(owner, repo, branch, fullPath string, headers *Headers) ([]Commit, error) {

	body, err := apis.HttpGet(gh.httpClient, APILastCommitsOfPath(owner, repo, branch, fullPath), headers.ToMap())
	if err != nil {
		return nil, err
	}

	var data []Commit
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// APIRepoTree github tree api
func APIRepoTree(owner, repo, branch string) string {
	return fmt.Sprintf("https://api.%s/repos/%s/%s/git/trees/%s?recursive=1", DEFAULT_HOST, owner, repo, branch)
}

// APIRaw github raw file api
func APIRaw(owner, repo, branch, path string) string {
	return fmt.Sprintf("https://%s/%s/%s/%s/%s", RAW_HOST, owner, repo, branch, path)
}

// APIDefaultBranch github repo metadata api
func APIMetadata(owner, repo string) string {
	return fmt.Sprintf("https://api.%s/repos/%s/%s", DEFAULT_HOST, owner, repo)
}

// APILastCommits github last commit api
func APILastCommits(owner, repo string) string {
	return fmt.Sprintf("https://api.%s/repos/%s/%s/commits", DEFAULT_HOST, owner, repo)
}

// APILastCommitsOfBranch github last commit of specific branch api
func APILastCommitsOfBranch(owner, repo, branch string) string {
	return fmt.Sprintf("%s/%s", APILastCommits(owner, repo), branch)
}

// APILastCommitsOfPath github last commit of specific branch api
func APILastCommitsOfPath(owner, repo, branch, path string) string {
	return fmt.Sprintf("%s?path=%s", APILastCommitsOfBranch(owner, repo, branch), path)
}
