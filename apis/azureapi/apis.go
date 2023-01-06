package azureapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kubescape/go-git-url/apis"
)

const (
	DEFAULT_HOST string = "azure.com"
	DEV_HOST     string = "dev.azure.com"
)

type IAzureAPI interface {
	GetRepoTree(owner, project, repo, branch string, headers *Headers) (*Tree, error)
	GetDefaultBranchName(owner, project, repo string, headers *Headers) (string, error)
	GetLatestCommit(owner, project, repo, branch string, headers *Headers) (*Commit, error)
}

type AzureAPI struct {
	httpClient *http.Client
}

func NewAzureAPI() *AzureAPI { return &AzureAPI{httpClient: &http.Client{}} }

func (az *AzureAPI) GetRepoTree(owner, project, repo, branch string, headers *Headers) (*Tree, error) {
	treeAPI := APIRepoTree(owner, project, repo, branch)
	body, err := apis.HttpGet(az.httpClient, treeAPI, headers.ToMap())
	if err != nil {
		return nil, err
	}

	var tree Tree
	if err = json.Unmarshal([]byte(body), &tree); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body from '%s', reason: %s", treeAPI, err.Error())
	}

	return &tree, nil

}

func (az *AzureAPI) GetDefaultBranchName(owner, project, repo string, headers *Headers) (string, error) {

	body, err := apis.HttpGet(az.httpClient, APIMetadata(owner, project, repo), headers.ToMap())
	if err != nil {
		return "", err
	}

	var data azureDefaultBranchAPI
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}
	for i := range data.BranchValue {
		if data.BranchValue[i].IsBaseVersion {
			return data.BranchValue[i].Name, nil
		}
	}

	errAPI := errors.New("unable to find default branch from the Azure API")
	return "", errAPI
}

func (az *AzureAPI) GetLatestCommit(owner, project, repo, branch string, headers *Headers) (*Commit, error) {

	body, err := apis.HttpGet(az.httpClient, APILastCommitsOfBranch(owner, project, repo, branch), headers.ToMap())
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
func (az *AzureAPI) GetFileLatestCommit(owner, project, repo, branch, fullPath string, headers *Headers) ([]Commit, error) {

	body, err := apis.HttpGet(az.httpClient, APILastCommitsOfPath(owner, project, repo, branch, fullPath), headers.ToMap())
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

// APIRepoTree Azure tree api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/git/items/list?view=azure-devops-rest-7.0&tabs=HTTP#full-recursion-and-with-content-metadata
// Example: https://dev.azure.com/anubhav06/testing/_apis/git/repositories/testing/items?recursionLevel=Full&versionDescriptor.version=dev&api-version=5.1
func APIRepoTree(owner, project, repo, branch string) string {
	return fmt.Sprintf("https://dev.%s/%s/%s/_apis/git/repositories/%s/items?recursionLevel=Full&versionDescriptor.version=%s&api-version=5.1", DEFAULT_HOST, owner, project, repo, branch)
}

// APIRaw Azure raw file api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/build/source-providers/get-file-contents?view=azure-devops-rest-7.0
// https://stackoverflow.com/questions/56281152/how-to-get-a-link-to-a-file-from-a-vso-repo/56283730#56283730
// Example: https://dev.azure.com/anubhav06/k8s-example/_apis/sourceProviders/tfsgit/filecontents?&repository=k8s-example&commitOrBranch=master&path=/volumes/cephfs/cephfs.yaml&api-version=7.0
func APIRaw(owner, project, repo, branch, path string) string {
	return fmt.Sprintf("https://dev.%s/%s/%s/_apis/sourceProviders/tfsgit/filecontents?&repository=%s&commitOrBranch=%s&path=%s", DEFAULT_HOST, owner, project, repo, branch, path)
}

// APIDefaultBranch Azure repo metadata api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/git/stats/list?view=azure-devops-rest-4.1&tabs=HTTP
// Example: https://dev.azure.com/anubhav06/k8s-example/_apis/git/repositories/k8s-example/stats/branches?api-version=4.1
func APIMetadata(owner, project, repo string) string {
	return fmt.Sprintf("https://dev.%s/%s/%s/_apis/git/repositories/%s/stats/branches?api-version=4.1", DEFAULT_HOST, owner, project, repo)
}

// APILastCommits Azure last commit api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/git/commits/get-commits?view=azure-devops-rest-4.1&tabs=HTTP
// Example: https://dev.azure.com/anubhav06/k8s-example/_apis/git/repositories/k8s-example/commits?searchCriteria.$top=1&api-version=4.1
func APILastCommits(owner, project, repo string) string {
	return fmt.Sprintf("https://dev.%s/%s/%s/_apis/git/repositories/%s/commits?searchCriteria.$top=1", DEFAULT_HOST, owner, project, repo)
}

// APILastCommitsOfBranch Azure last commit of specific branch api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/git/commits/get-commits?view=azure-devops-rest-4.1&tabs=HTTP
// Example: https://dev.azure.com/anubhav06/testing/_apis/git/repositories/testing/commits?searchCriteria.$top=1&searchCriteria.itemVersion.version=dev
func APILastCommitsOfBranch(owner, project, repo, branch string) string {
	return fmt.Sprintf("%s&searchCriteria.itemVersion.version=%s", APILastCommits(owner, project, repo), branch)
}

// APILastCommitsOfPath Azure last commit of specific branch api
// API Ref: https://learn.microsoft.com/en-us/rest/api/azure/devops/git/commits/get-commits?view=azure-devops-rest-4.1&tabs=HTTP
// Example: https://dev.azure.com/anubhav06/k8s-example/_apis/git/repositories/k8s-example/commits?searchCriteria.$top=1&searchCriteria.itemVersion.version=master&searchCriteria.itemPath=volumes/storageos/storageos-pod.yaml
func APILastCommitsOfPath(owner, project, repo, branch, path string) string {
	return fmt.Sprintf("%s&searchCriteria.itemPath=%s", APILastCommitsOfBranch(owner, project, repo, branch), path)
}
