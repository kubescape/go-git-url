package v1

import (
	"fmt"

	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/githubapi"
)

// ListAll list all directories and files in url tree
func (gh *GitHubURL) GetLatestCommit() (*apis.Commit, error) {
	if gh.GetHostName() == "" || gh.GetOwnerName() == "" || gh.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/repo")
	}
	if gh.GetBranchName() == "" {
		if err := gh.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	c, err := gh.gitHubAPI.GetLatestCommit(gh.GetOwnerName(), gh.GetRepoName(), gh.GetBranchName(), gh.headers())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit. reason: %s", err.Error())
	}

	return gitHubAPICommitToCommit(c), nil
}

func gitHubAPICommitToCommit(c *githubapi.Commit) *apis.Commit {
	latestCommit := &apis.Commit{
		SHA: c.SHA,
		Author: apis.Committer{
			Name:  c.Commit.Author.Name,
			Email: c.Commit.Author.Email,
			Date:  c.Commit.Author.Date,
		},
		Committer: apis.Committer{
			Name:  c.Commit.Committer.Name,
			Email: c.Commit.Committer.Email,
			Date:  c.Commit.Committer.Date,
		},
		Message: c.Commit.Message,
	}
	for i := range c.Files {
		latestCommit.Files = append(latestCommit.Files, apis.Files{
			FileSHA:  c.Files[i].SHA,
			Filename: c.Files[i].Filename,
		})

	}
	return latestCommit
}
