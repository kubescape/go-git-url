package v1

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/armosec/url-git-go/apis/githubapi"

	"k8s.io/utils/strings/slices"
)

// NewGitHubParser empty instance of a github parser
func NewGitHubParser() *GitHubURL {
	return &GitHubURL{
		host:  "github.com",
		token: os.Getenv("GITHUB_TOKEN"),
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

func (gh *GitHubURL) GetHost() string   { return gh.host }
func (gh *GitHubURL) GetBranch() string { return gh.branch }
func (gh *GitHubURL) GetOwner() string  { return gh.owner }
func (gh *GitHubURL) GetRepo() string   { return gh.repo }
func (gh *GitHubURL) GetPath() string   { return gh.path }
func (gh *GitHubURL) GetToken() string  { return gh.token }

func (gh *GitHubURL) SetOwner(o string)       { gh.owner = o }
func (gh *GitHubURL) SetRepo(r string)        { gh.repo = r }
func (gh *GitHubURL) SetPath(p string)        { gh.path = p }
func (gh *GitHubURL) SetBranch(branch string) { gh.branch = branch }
func (gh *GitHubURL) SetToken(token string)   { gh.token = token }

// Parse URL
func (gh *GitHubURL) Parse(fullURL string) error {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return err
	}

	if parsedURL.Host == "raw.githubusercontent.com" {
		gh.isFile = true
	}

	index := 0

	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' })
	if len(splittedRepo) < 2 {
		return fmt.Errorf("expecting <user>/<repo> in url path, received: '%s'", parsedURL.Path)
	}
	gh.owner = splittedRepo[index]
	index += 1
	gh.repo = splittedRepo[index]
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
func (gh *GitHubURL) SetDefaultBranch() error {
	branch, err := githubapi.GetDefaultBranchName(gh.GetOwner(), gh.GetRepo(), gh.headres())
	if err != nil {
		return err
	}
	gh.branch = branch
	return nil
}

func (gh *GitHubURL) ListFiles(filesExtensions []string) ([]string, error) {
	if gh.GetHost() == "" || gh.GetOwner() == "" || gh.GetRepo() == "" {
		return []string{}, fmt.Errorf("missing host/owner/repo")
	}
	if gh.GetBranch() == "" {
		if err := gh.SetDefaultBranch(); err != nil {
			return []string{}, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	// if the URL points directly to a file
	if gh.isFile {
		if slices.Contains(filesExtensions, getFileExtension(gh.GetPath())) {
			return []string{githubapi.APIRaw(gh.GetOwner(), gh.GetRepo(), gh.GetBranch(), gh.GetPath())}, nil
		} else {
			return []string{}, nil
		}
	}

	repoTree, err := githubapi.GetRepoTree(gh.GetOwner(), gh.GetRepo(), gh.GetBranch(), gh.headres())
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	var files []string
	for _, path := range repoTree.ListAll() {
		if gh.path != "" && !strings.HasPrefix(path, gh.GetPath()) {
			continue
		}
		if len(filesExtensions) > 0 {
			if slices.Contains(filesExtensions, getFileExtension(path)) {
				files = append(files, githubapi.APIRaw(gh.GetOwner(), gh.GetRepo(), gh.GetBranch(), gh.GetPath()))
			}
		} else {
			files = append(files, githubapi.APIRaw(gh.GetOwner(), gh.GetRepo(), gh.GetBranch(), gh.GetPath()))
		}
	}
	return files, nil
}

func getFileExtension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}
func (gh *GitHubURL) headres() *githubapi.Headres {
	return &githubapi.Headres{Token: gh.GetToken()}
}
