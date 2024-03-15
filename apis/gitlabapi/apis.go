package gitlabapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kubescape/go-git-url/apis"
)

type IGitLabAPI interface {
	GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error)
	GetDefaultBranchName(owner, repo string, headers *Headers) (string, error)
	GetLatestCommit(owner, repo, branch string, headers *Headers) (*Commit, error)
}

type GitLabAPI struct {
	host       string
	httpClient *http.Client
}

func NewGitLabAPI(host string) *GitLabAPI {
	return &GitLabAPI{
		host:       host,
		httpClient: &http.Client{},
	}
}

func getProjectId(owner, repo string) string {
	return strings.ReplaceAll(owner+"/"+repo, "/", "%2F")
}

func (gl *GitLabAPI) GetRepoTree(owner, repo, branch string, headers *Headers) (*Tree, error) {
	id := getProjectId(owner, repo)

	treeAPI := APIRepoTree(gl.host, id, branch)
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
	id := getProjectId(owner, repo)

	body, err := apis.HttpGet(gl.httpClient, APIMetadata(gl.host, id), headers.ToMap())
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
	id := getProjectId(owner, repo)

	body, err := apis.HttpGet(gl.httpClient, APILastCommitsOfBranch(gl.host, id, branch), headers.ToMap())
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
func APIRepoTree(host, id, branch string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/tree?ref=%s", host, id, branch)
}

// APIRaw GitLab : Get raw file from repository
// API Ref: https://docs.gitlab.com/ee/api/repository_files.html#get-raw-file-from-repository
// Example: https://gitlab.com/api/v4/projects/23383112/repository/files/app%2Findex.html/raw
func APIRaw(host, owner, repo, path string) string {
	id := getProjectId(owner, repo)

	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/files/%s/raw", host, id, path)
}

// APIDefaultBranch GitLab : Branch metadata; list repo branches
// API Ref: https://docs.gitlab.com/ee/api/branches.html#list-repository-branches
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/branches
func APIMetadata(host, id string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/branches", host, id)
}

// APILastCommits GitLab : List repository commits
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits
func APILastCommits(host, id string) string {
	return fmt.Sprintf("https://%s/api/v4/projects/%s/repository/commits", host, id)
}

// TODO: These return list of commits, and we might want just a single latest commit. Pls check with this later:

// APILastCommitsOfBranch GitLab : Last commits of specific branch
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits?ref_name=feature/k8s-in-hour
func APILastCommitsOfBranch(host, id, branch string) string {
	return fmt.Sprintf("%s?ref_name=%s", APILastCommits(host, id), branch)
}

// APILastCommitsOfPath GitLab : Last commits of specific branch & specified path
// API Ref: https://docs.gitlab.com/ee/api/commits.html#list-repository-commits
// Example: https://gitlab.com/api/v4/projects/nanuchi%2Fdeveloping-with-docker/repository/commits?ref_name=master&path=app/server.js
func APILastCommitsOfPath(host, id, branch, path string) string {
	return fmt.Sprintf("%s&path=%s", APILastCommitsOfBranch(host, id, branch), path)
}
