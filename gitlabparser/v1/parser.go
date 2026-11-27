package v1

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	giturl "github.com/chainguard-dev/git-urls"
	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/gitlabapi"
)

// NewGitLabParserWithURL parses a repository on the configured host (including any
// non-default port). An empty host selects gitlab.com.
func NewGitLabParserWithURL(host, fullURL string) (*GitLabURL, error) {
	gl := &GitLabURL{
		configuredHost: host,
		token:          os.Getenv("GITLAB_TOKEN"),
	}

	if err := gl.Parse(fullURL); err != nil {
		return nil, err
	}

	return gl, nil
}

func (gl *GitLabURL) GetURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   gl.GetHostName(),
		Path:   fmt.Sprintf("%s/%s", gl.GetOwnerName(), gl.GetRepoName()),
	}
}

func IsHostGitLab(host string) bool {
	return strings.EqualFold(host, "gitlab.com")
}

func (gl *GitLabURL) GetProvider() string    { return apis.ProviderGitLab.String() }
func (gl *GitLabURL) GetHostName() string    { return gl.host }
func (gl *GitLabURL) GetProjectName() string { return gl.project }
func (gl *GitLabURL) GetBranchName() string  { return gl.branch }
func (gl *GitLabURL) GetOwnerName() string   { return gl.owner }
func (gl *GitLabURL) GetRepoName() string    { return gl.repo }
func (gl *GitLabURL) GetPath() string        { return gl.path }
func (gl *GitLabURL) GetToken() string       { return gl.token }
func (gl *GitLabURL) GetHttpCloneURL() string {
	return fmt.Sprintf("https://%s/%s/%s.git", gl.host, gl.owner, gl.repo)
}

func (gl *GitLabURL) SetOwnerName(o string)         { gl.owner = o }
func (gl *GitLabURL) SetProjectName(project string) { gl.project = project }
func (gl *GitLabURL) SetRepoName(r string)          { gl.repo = r }
func (gl *GitLabURL) SetBranchName(branch string)   { gl.branch = branch }
func (gl *GitLabURL) SetPath(p string)              { gl.path = p }
func (gl *GitLabURL) SetToken(token string)         { gl.token = token }

// Parse URL
func (gl *GitLabURL) Parse(fullURL string) error {
	parsedURL, err := giturl.Parse(fullURL)
	if err != nil {
		return err
	}

	host := apis.HostWithPort(parsedURL)
	allowed := IsHostGitLab(parsedURL.Hostname())
	if gl.configuredHost != "" {
		allowed = strings.EqualFold(host, gl.configuredHost)
	}
	if !allowed {
		return fmt.Errorf("repository host '%s' not supported", host)
	}

	index := 0

	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice
	if len(splittedRepo) < 1 {
		return fmt.Errorf("expecting <user>/<repo> or <repo> in url path, received: '%s'", parsedURL.Path)
	}

	if gl.configuredHost != "" {
		host = gl.configuredHost
	}
	gl.host = host
	gl.gitLabAPI = gitlabapi.NewGitLabAPI(host)

	// in gitlab <user>/<repo> are separated from blob/tree/raw with a -
	for i := range splittedRepo {
		if splittedRepo[i] == "-" {
			break
		}
		index = i
	}

	gl.owner = strings.Join(splittedRepo[:index], "/")
	gl.repo = strings.TrimSuffix(splittedRepo[index], ".git")
	index++

	// root of repo
	if len(splittedRepo) < index+1 {
		return nil
	}

	if splittedRepo[index] == "-" {
		index++ // skip "-" symbol in URL
	}

	if len(splittedRepo) < index+1 {
		return nil
	}

	// is file or dir
	switch splittedRepo[index] {
	case "blob", "tree", "raw":
		index++
	}

	if len(splittedRepo) < index+1 {
		return nil
	}

	gl.branch = splittedRepo[index]
	index += 1

	if len(splittedRepo) < index+1 {
		return nil
	}
	gl.path = strings.Join(splittedRepo[index:], "/")

	return nil
}

// Set the default brach of the repo
func (gl *GitLabURL) SetDefaultBranchName() error {
	branch, err := gl.gitLabAPI.GetDefaultBranchName(gl.GetOwnerName(), gl.GetRepoName(), gl.headers())
	if err != nil {
		return err
	}
	gl.branch = branch
	return nil
}

func (gl *GitLabURL) headers() *gitlabapi.Headers {
	return &gitlabapi.Headers{Token: gl.GetToken()}
}
