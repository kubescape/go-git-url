package v1

import (
	"fmt"

	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/gitlabapi"
)

// ListAll list all directories and files in url tree
func (gl *GitLabURL) GetLatestCommit() (*apis.Commit, error) {
	if gl.GetHostName() == "" || gl.GetOwnerName() == "" || gl.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/repo")
	}
	if gl.GetBranchName() == "" {
		if err := gl.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	c, err := gl.gitLabAPI.GetLatestCommit(gl.GetOwnerName(), gl.GetRepoName(), gl.GetBranchName(), gl.headers())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit. reason: %s", err.Error())
	}

	return gitLabAPICommitToCommit(c), nil
}

func gitLabAPICommitToCommit(c *gitlabapi.Commit) *apis.Commit {
	latestCommit := &apis.Commit{
		SHA: c.ID,
		Author: apis.Committer{
			Name:  c.AuthorName,
			Email: c.AuthorEmail,
			Date:  c.AuthoredDate,
		},
		Committer: apis.Committer{
			Name:  c.CommitterName,
			Email: c.CommitterEmail,
			Date:  c.CommitterDate,
		},
		Message: c.Message,
	}

	return latestCommit
}
