package giturl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitURL(t *testing.T) {
	fileText := "https://raw.githubusercontent.com/armosec/go-git-url/master/files/file0.text"
	var gitURL IGitURL
	var err error
	{
		gitURL, err = NewGitURL("https://github.com/armosec/go-git-url")
		assert.NoError(t, err)

		files, err := gitURL.ListFilesNamesWithExtension([]string{"yaml", "json"})
		assert.NoError(t, err)
		assert.Equal(t, 7, len(files))
	}

	{
		gitURL, err = NewGitURL("https://github.com/armosec/go-git-url")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitURL(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err = NewGitURL(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadAllFiles()
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err := NewGitURL("https://github.com/armosec/go-git-url/tree/master/files")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitURL("https://github.com/armosec/go-git-url/blob/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}
}
