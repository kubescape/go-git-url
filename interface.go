package main

import "net/url"

// IGitURL parse git urls
type IGitURL interface {

	// parse url
	Parse(fullURL string) error

	SetBranch(string)
	SetOwner(string)
	SetPath(string)
	SetToken(string)

	GetBranch() string
	GetOwner() string
	GetPath() string
	GetToken() string

	// GetURL git url
	GetURL() *url.URL

	// ListFilesNamesWithExtension list all files in path with the desired extension. if empty will list all (including directories)
	ListFilesNamesWithExtension(exctensions []string) ([]string, error)

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
	DownloadFilesWithExtension(exctensions []string) (map[string][]byte, map[string]error)
}
