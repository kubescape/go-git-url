package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/armosec/url-git-go/apis"
)

func APIRepoTree(owner, repo, branch string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", owner, repo, branch)
}

func APIRaw(owner, repo, branch, path string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", owner, repo, branch, path)
}
func APIDefaultBranch(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
}

func GetRepoTree(owner, repo, branch string, headres *Headres) (*Tree, error) {
	treeAPI := APIRepoTree(owner, repo, branch)
	body, err := apis.HttpGet(&http.Client{}, treeAPI, headres.ToMap())
	if err != nil {
		return nil, err
	}

	var tree Tree
	err = json.Unmarshal([]byte(body), &tree)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", treeAPI, err.Error())
		// return nil
	}
	return &tree, nil

}

func GetDefaultBranchName(owner, repo string, headres *Headres) (string, error) {

	body, err := apis.HttpGet(&http.Client{}, APIDefaultBranch(owner, repo), headres.ToMap())
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
