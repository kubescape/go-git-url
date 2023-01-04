package v1

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/bitbucketapi"
)

var rawUserRe = regexp.MustCompile("([^<]*)?(<(.+)>)?")

func (gl *BitBucketURL) GetLatestCommit() (*apis.Commit, error) {
	if gl.GetHostName() == "" || gl.GetOwnerName() == "" || gl.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/repo")
	}
	if gl.GetBranchName() == "" {
		if err := gl.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	c, err := gl.bitBucketAPI.GetLatestCommit(gl.GetOwnerName(), gl.GetRepoName(), gl.GetBranchName(), gl.headers())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit. reason: %s", err.Error())
	}

	return bitBucketAPICommitToCommit(c), nil
}

func bitBucketAPICommitToCommit(c *bitbucketapi.Commit) *apis.Commit {
	name, email := parseRawUser(c.Author.Raw)
	latestCommit := &apis.Commit{
		SHA: c.Hash,
		Author: apis.Committer{
			Name:  name,
			Email: email,
			Date:  c.Date,
		},
		Committer: apis.Committer{ // same as author as API doesn't return the committer
			Name:  name,
			Email: email,
			Date:  c.Date,
		},
		Message: c.Message,
	}

	return latestCommit
}

func parseRawUser(raw string) (string, string) {
	match := rawUserRe.FindStringSubmatch(raw)
	if match != nil {
		return strings.TrimSpace(match[1]), match[3]
	}
	return raw, ""
}
