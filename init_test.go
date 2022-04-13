package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitURL(t *testing.T) {
	{
		gitURL, err := NewGitURL("https://github.com/armosec/url-git-go")
		assert.NoError(t, err)

		files, err := gitURL.ListFilesNamesWithExtension([]string{"yaml", "json"})
		assert.NoError(t, err)
		assert.Equal(t, 6, len(files))
	}

	{
		gitURL, err := NewGitURL("https://github.com/armosec/url-git-go")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files["https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text"]))

	}

	{
		gitURL, err := NewGitURL("https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files["https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text"]))
	}

	{
		gitURL, err := NewGitURL("https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadAllFiles()
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files["https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text"]))
	}

	{
		gitURL, err := NewGitURL("https://github.com/armosec/url-git-go/tree/master/files")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files["https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text"]))

	}

	{
		gitURL, err := NewGitURL("https://github.com/armosec/url-git-go/blob/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files["https://raw.githubusercontent.com/armosec/url-git-go/master/files/file0.text"]))

	}
}
