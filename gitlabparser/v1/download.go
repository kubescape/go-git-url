package v1

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/kubescape/go-git-url/apis/gitlabapi"
)

// DownloadAllFiles download files from git repo tree
// return map of (url:file, url:error)
func (gl *GitLabURL) DownloadAllFiles() (map[string][]byte, map[string]error) {

	files, err := gl.ListFilesNames()
	if err != nil {
		return nil, map[string]error{"": err} // TODO - update error
	}
	return downloadFiles(gl.pathsToURLs(files))
}

// DownloadFilesWithExtension download files from git repo tree based on file extension
// return map of (url:file, url:error)
func (gl *GitLabURL) DownloadFilesWithExtension(filesExtensions []string) (map[string][]byte, map[string]error) {
	files, err := gl.ListFilesNamesWithExtension(filesExtensions)
	if err != nil {
		return nil, map[string]error{"": err} // TODO - update error
	}

	return downloadFiles(gl.pathsToURLs(files))
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

func (gl *GitLabURL) pathsToURLs(files []string) []string {
	var rawURLs []string
	for i := range files {
		rawURLs = append(rawURLs, gitlabapi.APIRaw(gl.host, gl.GetOwnerName(), gl.GetRepoName(), files[i]))
	}
	return rawURLs
}
