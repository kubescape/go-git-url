package v1

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/armosec/go-git-url/apis/githubapi"
)

// DownloadAllFiles download files from git repo tree
// return map of (url:file, url:error)
func (gh *GitHubURL) DownloadAllFiles() (map[string][]byte, map[string]error) {
	files, err := gh.ListFilesNames()
	if err != nil {
		return nil, map[string]error{"": err} // TODO - update error
	}
	return downloadFiles(gh.pathsToURLs(files))
}

// DownloadFilesWithExtension download files from git repo tree based on file extension
// return map of (url:file, url:error)
func (gh *GitHubURL) DownloadFilesWithExtension(filesExtensions []string) (map[string][]byte, map[string]error) {
	files, err := gh.ListFilesNamesWithExtension(filesExtensions)
	if err != nil {
		return nil, map[string]error{"": err} // TODO - update error
	}

	return downloadFiles(gh.pathsToURLs(files))
}

// DownloadFiles download files from git repo. Call ListAllNames/ListFilesNamesWithExtention before calling this function
// return map of (url:file, url:error)
func downloadFiles(urls []string) (map[string][]byte, map[string]error) {
	var errs map[string]error
	files := map[string][]byte{}
	for i := range urls {

		file, err := downloadFile(urls[i])
		if err != nil {
			if errs == nil {
				errs = map[string]error{}
			}
			errs[urls[i]] = err
			continue
		}
		files[urls[i]] = file
	}
	return files, errs
}

// download single file
func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || 301 < resp.StatusCode {
		return nil, fmt.Errorf("failed to download file, url: '%s', status code: %s", url, resp.Status)
	}
	return streamToByte(resp.Body)
}

func streamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stream); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (gh *GitHubURL) pathsToURLs(files []string) []string {
	var rawURLs []string
	for i := range files {
		rawURLs = append(rawURLs, githubapi.APIRaw(gh.GetOwnerName(), gh.GetRepoName(), gh.GetBranchName(), files[i]))
	}
	return rawURLs
}
