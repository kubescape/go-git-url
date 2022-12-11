package gitlabapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kubescape/go-git-url/apis"
)

const (
	DEFAULT_HOST string = "gitlab.com"
)

type IGitLabAPI interface {
	GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error)
	GetDefaultBranchName(owner, repo string, headers *Headers) (string, error)
	GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error)
}

type GitLabAPI struct {
	httpClient *http.Client
}

func NewGitLabAPI() *GitLabAPI { return &GitLabAPI{httpClient: &http.Client{}} }

func (gl *GitLabAPI) GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error) {
	id := owner + "%2F" + repo

	treeAPI := APIRepoTree(id, branch)
	body, err := apis.HttpGet(gl.httpClient, treeAPI, headers.ToMap())
	if err != nil {
		return nil, err
	}

	var tree Tree
	if err = json.Unmarshal([]byte(body), &tree); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", treeAPI, err.Error())
	}

	return &tree, nil

}

func (gl *GitLabAPI) GetDefaultBranchName(owner, repo string, headers *Headers) (string, error) {

	id := owner + "%2F" + repo
	body, err := apis.HttpGet(gl.httpClient, APIMetadata(id), headers.ToMap())
	if err != nil {
		return "", err
	}

	var data gitlabDefaultBranchAPI
	if err = json.Unmarshal([]byte(body), &data); err != nil {
		return "", err
	}

	for i := range data {
		if data[i].DefaultBranch {
			return data[i].Name, nil
		}
	}

	errAPI := errors.New("unable to find default branch from the GitLab API")
	return "", errAPI
}

func (gl *GitLabAPI) GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error) {

	id := owner + "%2F" + repo

	body, err := apis.HttpGet(gl.httpClient, APILastCommitsOfBranch(id, branch), headers.ToMap())
	if err != nil {
		return nil, err
	}

	var data []Commit
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return &data[0], err
	}
	return &data[0], nil
}

// APIRepoTree GitLab tree api
// API Ref: https://docs.gitlab.com/ee/api/repositories.html
func APIRepoTree(id, branch string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/tree?ref=%s", DEFAULT_HOST, id, branch)
}

// APIRaw GitLab : Get raw file from repository
// API Ref: https://docs.gitlab.com/ee/api/repository_files.html#get-raw-file-from-repository
// Example: https://gitlab.com/api/v4/projects/23383112/repository/files/app%2Findex.html/raw
func APIRaw(owner, repo, branch, path string) string {
	id := owner + "%2F" + repo
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/files/%s/raw", DEFAULT_HOST, id, path)
}

// APIDefaultBranch GitLab : Branch metadata; list repo branches
// API Ref: https://docs.gitlab.com/ee/api/branches.html#list-repository-branches
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/branches
func APIMetadata(id string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/branches", DEFAULT_HOST, id)
}

// APILastCommits GitLab : List repository commits
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits
func APILastCommits(id string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/commits", DEFAULT_HOST, id)
}

// TODO: These return list of commits, and we might want just a single latest commit. Pls check with this later:

// APILastCommitsOfBranch GitLab : Last commits of specific branch
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits?ref_name=feature/k8s-in-hour
func APILastCommitsOfBranch(id, branch string) string {
	return fmt.Sprintf("%s?ref_name=%s", APILastCommits(id), branch)
}

// APILastCommitsOfPath GitLab : Last commits of specific branch & specified path
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits?ref_name=master&path=app/server.js
func APILastCommitsOfPath(id, branch, path string) string {
	return fmt.Sprintf("%s&path=%s", APILastCommitsOfBranch(id, branch), path)
}
