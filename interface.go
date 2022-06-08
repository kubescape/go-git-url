package giturl

import (
	"net/url"

	"github.com/armosec/go-git-url/apis"
)

// IGitURL parse git urls
type IGitURL interface {
	SetBranchName(string)
	SetOwnerName(string)
	SetPath(string)
	SetRepoName(string)

	GetProvider() string
	GetBranchName() string
	GetOwnerName() string
	GetPath() string
	GetRepoName() string

	// parse url
	Parse(fullURL string) error

	// GetURL git url
	GetURL() *url.URL
}

type IGitAPI interface {
	IGitURL

	GetToken() string
	SetToken(string)

	// set default branch name using the providers git API
	SetDefaultBranchName() error

	// ListFilesNamesWithExtension list all files in path with the desired extension. if empty will list all (including directories)
	ListFilesNamesWithExtension(extensions []string) ([]string, error)

	// ListAll list all directories and files in url tree
	ListAllNames() ([]string, error)

	// ListFilesNames list all files in url tree
	ListFilesNames() ([]string, error)

	// ListDirsNames list all directories in url tree
	ListDirsNames() ([]string, error)

	// DownloadAllFiles download files from git repo tree
	// return map of (url:file, url:error)
	DownloadAllFiles() (map[string][]byte, map[string]error)

	// DownloadFilesWithExtension download files from git repo tree based on file extension
	// return map of (url:file, url:error)
	DownloadFilesWithExtension(extensions []string) (map[string][]byte, map[string]error)

	// GetLatestCommit get latest commit
	GetLatestCommit() (*apis.Commit, error)
}
