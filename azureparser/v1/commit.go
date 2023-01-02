package v1

import (
	"fmt"

	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/azureapi"
)

// ListAll list all directories and files in url tree
func (az *AzureURL) GetLatestCommit() (*apis.Commit, error) {
	if az.GetHostName() == "" || az.GetOwnerName() == "" || az.GetProjectName() == "" || az.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/project/repo")
	}
	if az.GetBranchName() == "" {
		if err := az.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	c, err := az.azureAPI.GetLatestCommit(az.GetOwnerName(), az.GetProjectName(), az.GetRepoName(), az.GetBranchName(), az.headers())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit. reason: %s", err.Error())
	}

	return azureAPICommitToCommit(c), nil
}

func azureAPICommitToCommit(c *azureapi.Commit) *apis.Commit {
	latestCommit := &apis.Commit{
		SHA: c.CommitsValue[0].CommitID,
		Author: apis.Committer{
			Name:  c.CommitsValue[0].Author.Name,
			Email: c.CommitsValue[0].Author.Email,
			Date:  c.CommitsValue[0].Author.Date,
		},
		Committer: apis.Committer{
			Name:  c.CommitsValue[0].Committer.Name,
			Email: c.CommitsValue[0].Committer.Email,
			Date:  c.CommitsValue[0].Committer.Date,
		},
		Message: c.CommitsValue[0].Comment,
	}

	return latestCommit
}
