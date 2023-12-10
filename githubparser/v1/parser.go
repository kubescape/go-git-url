package v1

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	giturl "github.com/chainguard-dev/git-urls"
	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/githubapi"
)

// NewGitHubParser empty instance of a github parser
func NewGitHubParser() *GitHubURL {

	return &GitHubURL{
		gitHubAPI: githubapi.NewGitHubAPI(),
		host:      githubapi.DEFAULT_HOST,
		token:     os.Getenv("GITHUB_TOKEN"),
	}
}

// NewGitHubParserWithURL parsed instance of a github parser
func NewGitHubParserWithURL(fullURL string) (*GitHubURL, error) {
	gh := NewGitHubParser()

	if err := gh.Parse(fullURL); err != nil {
		return gh, err
	}

	return gh, nil
}

func (gh *GitHubURL) GetURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   gh.GetHostName(),
		Path:   fmt.Sprintf("%s/%s", gh.GetOwnerName(), gh.GetRepoName()),
	}
}
func IsHostGitHub(host string) bool {
	return host == githubapi.DEFAULT_HOST || host == githubapi.RAW_HOST || host == githubapi.SUBDOMAIN_HOST
}

func (gh *GitHubURL) GetProvider() string   { return apis.ProviderGitHub.String() }
func (gh *GitHubURL) GetHostName() string   { return gh.host }
func (gh *GitHubURL) GetBranchName() string { return gh.branch }
func (gh *GitHubURL) GetOwnerName() string  { return gh.owner }
func (gh *GitHubURL) GetRepoName() string   { return gh.repo }
func (gh *GitHubURL) GetPath() string       { return gh.path }
func (gh *GitHubURL) GetToken() string      { return gh.token }
func (gh *GitHubURL) GetHttpCloneURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", gh.GetOwnerName(), gh.GetRepoName())
}

func (gh *GitHubURL) SetOwnerName(o string)       { gh.owner = o }
func (gh *GitHubURL) SetRepoName(r string)        { gh.repo = r }
func (gh *GitHubURL) SetBranchName(branch string) { gh.branch = branch }
func (gh *GitHubURL) SetPath(p string)            { gh.path = p }
func (gh *GitHubURL) SetToken(token string)       { gh.token = token }

// Parse URL
func (gh *GitHubURL) Parse(fullURL string) error {
	parsedURL, err := giturl.Parse(fullURL)
	if err != nil {
		return err
	}

	if parsedURL.Host == githubapi.RAW_HOST {
		gh.isFile = true
	}

	index := 0

	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice
	if len(splittedRepo) < 2 {
		return fmt.Errorf("expecting <user>/<repo> in url path, received: '%s'", parsedURL.Path)
	}
	gh.owner = splittedRepo[index]
	index += 1
	gh.repo = strings.TrimSuffix(splittedRepo[index], ".git")
	index += 1

	// root of repo
	if len(splittedRepo) < index+1 {
		return nil
	}

	// is file or dir
	switch splittedRepo[index] {
	case "blob":
		gh.isFile = true
		index += 1
	case "tree":
		gh.isFile = false
		index += 1
	}

	if len(splittedRepo) < index+1 {
		return nil
	}

	gh.branch = splittedRepo[index]
	index += 1

	if len(splittedRepo) < index+1 {
		return nil
	}
	gh.path = strings.Join(splittedRepo[index:], "/")

	return nil
}

// Set the default brach of the repo
func (gh *GitHubURL) SetDefaultBranchName() error {
	branch, err := gh.gitHubAPI.GetDefaultBranchName(gh.GetOwnerName(), gh.GetRepoName(), gh.headers())
	if err != nil {
		return err
	}
	gh.branch = branch
	return nil
}

func (gh *GitHubURL) headers() *githubapi.Headers {
	return &githubapi.Headers{Token: gh.GetToken()}
}
