package apis

import (
	"fmt"
)

type UrlComposer interface {
	FileUrlByCommit(commit string) string
	FileUrlByBranch(branch string) string
	FileUrlByTag(tag string) string
}

type GitHubUrlComposer struct {
	remoteUrl, path string
}

type AzureUrlComposer struct {
	remoteUrl, path string
}

func NewUrlComposer(provider ProviderType, remoteUrl, path string) (UrlComposer, error) {
	switch provider {
	case ProviderAzure:
		return &AzureUrlComposer{remoteUrl: remoteUrl, path: path}, nil
	case ProviderGitHub:
		return &GitHubUrlComposer{remoteUrl: remoteUrl, path: path}, nil
	default:
		return nil, fmt.Errorf("unknown provider")
	}
}

func (r *GitHubUrlComposer) fileUrlBase(commitOrBranchOrTag string) string {
	return fmt.Sprintf("%s/blob/%s/%s", r.remoteUrl, commitOrBranchOrTag, r.path)
}

func (r *GitHubUrlComposer) FileUrlByCommit(commit string) string {
	return r.fileUrlBase(commit)
}

func (r *GitHubUrlComposer) FileUrlByBranch(branch string) string {
	return r.fileUrlBase(branch)
}

func (r *GitHubUrlComposer) FileUrlByTag(tag string) string {
	return r.fileUrlBase(tag)
}

func (r *AzureUrlComposer) fileUrlBase(version string) string {
	return fmt.Sprintf("%s?path=%s&version=%s", r.remoteUrl, r.path, version)
}

func (r *AzureUrlComposer) FileUrlByCommit(commit string) string {
	return r.fileUrlBase(fmt.Sprintf("GC%s", commit))
}

func (r *AzureUrlComposer) FileUrlByBranch(branch string) string {
	return r.fileUrlBase(fmt.Sprintf("GB%s", branch))
}

func (r *AzureUrlComposer) FileUrlByTag(tag string) string {
	return r.fileUrlBase(fmt.Sprintf("GT%s", tag))
}
