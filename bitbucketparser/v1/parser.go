package v1

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	giturl "github.com/chainguard-dev/git-urls"
	"github.com/kubescape/go-git-url/apis"
	"github.com/kubescape/go-git-url/apis/bitbucketapi"
)

const HOST = "bitbucket.org"

// NewBitBucketParser empty instance of a bitbucket parser
func NewBitBucketParser() *BitBucketURL {

	return &BitBucketURL{
		bitBucketAPI: bitbucketapi.NewBitBucketAPI(),
		host:         HOST,
		token:        os.Getenv("BITBUCKET_TOKEN"),
	}
}

// NewBitBucketParserWithURL parsed instance of a bitbucket parser
func NewBitBucketParserWithURL(fullURL string) (*BitBucketURL, error) {
	gl := NewBitBucketParser()

	if err := gl.Parse(fullURL); err != nil {
		return gl, err
	}

	return gl, nil
}

func (gl *BitBucketURL) GetURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   gl.GetHostName(),
		Path:   fmt.Sprintf("%s/%s", gl.GetOwnerName(), gl.GetRepoName()),
	}
}

func IsHostBitBucket(host string) bool { return strings.HasSuffix(host, HOST) }

func (gl *BitBucketURL) GetProvider() string    { return apis.ProviderBitBucket.String() }
func (gl *BitBucketURL) GetHostName() string    { return gl.host }
func (gl *BitBucketURL) GetProjectName() string { return gl.project }
func (gl *BitBucketURL) GetBranchName() string  { return gl.branch }
func (gl *BitBucketURL) GetOwnerName() string   { return gl.owner }
func (gl *BitBucketURL) GetRepoName() string    { return gl.repo }
func (gl *BitBucketURL) GetPath() string        { return gl.path }
func (gl *BitBucketURL) GetToken() string       { return gl.token }
func (gl *BitBucketURL) GetHttpCloneURL() string {
	return fmt.Sprintf("https://bitbucket.org/%s/%s.git", gl.GetOwnerName(), gl.GetRepoName())
}

func (gl *BitBucketURL) SetOwnerName(o string)         { gl.owner = o }
func (gl *BitBucketURL) SetProjectName(project string) { gl.project = project }
func (gl *BitBucketURL) SetRepoName(r string)          { gl.repo = r }
func (gl *BitBucketURL) SetBranchName(branch string)   { gl.branch = branch }
func (gl *BitBucketURL) SetPath(p string)              { gl.path = p }
func (gl *BitBucketURL) SetToken(token string)         { gl.token = token }

// Parse URL
func (gl *BitBucketURL) Parse(fullURL string) error {
	parsedURL, err := giturl.Parse(fullURL)
	if err != nil {
		return err
	}

	index := 0

	splitRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice
	if len(splitRepo) < 2 {
		return fmt.Errorf("expecting <user>/<repo> in url path, received: '%s'", parsedURL.Path)
	}
	gl.owner = splitRepo[index]
	index++
	gl.repo = strings.TrimSuffix(splitRepo[index], ".git")
	index++

	// root of repo
	if len(splitRepo) < index+1 {
		return nil
	}

	if splitRepo[index] == "-" {
		index++ // skip "-" symbol in URL
	}

	// is file or dir
	switch splitRepo[index] {
	case "src", "raw":
		index++
	}

	if len(splitRepo) < index+1 {
		return nil
	}

	gl.branch = splitRepo[index]
	index += 1

	if len(splitRepo) < index+1 {
		return nil
	}
	gl.path = strings.Join(splitRepo[index:], "/")

	return nil
}

// SetDefaultBranchName sets the default brach of the repo
func (gl *BitBucketURL) SetDefaultBranchName() error {
	branch, err := gl.bitBucketAPI.GetDefaultBranchName(gl.GetOwnerName(), gl.GetRepoName(), gl.headers())
	if err != nil {
		return err
	}
	gl.branch = branch
	return nil
}

func (gl *BitBucketURL) headers() *bitbucketapi.Headers {
	return &bitbucketapi.Headers{Token: gl.GetToken()}
}
