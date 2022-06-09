package v1

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	giturl "github.com/whilp/git-urls"
)

const HOST = "azure.com"
const HOST_DEV = "dev.azure.com"
const HOST_PROD = "prod.azure.com"

// NewGitHubParser empty instance of a github parser
func NewAzureParser() *AzureURL {

	return &AzureURL{
		token: os.Getenv("AZURE_TOKEN"),
	}
}

// NewGitHubParserWithURL parsed instance of a github parser
func NewAzureParserWithURL(fullURL string) (*AzureURL, error) {
	az := NewAzureParser()

	if err := az.Parse(fullURL); err != nil {
		return az, err
	}

	return az, nil
}

func (az *AzureURL) GetURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   az.host,
		Path:   fmt.Sprintf("%s/%s/_git/%s", az.GetOwnerName(), az.GetProjectName(), az.GetRepoName()),
	}
}

func IsHostAzure(host string) bool { return strings.HasSuffix(host, HOST) }

func (az *AzureURL) GetProvider() string    { return "azure" }
func (az *AzureURL) GetHostName() string    { return az.host }
func (az *AzureURL) GetProjectName() string { return az.project }
func (az *AzureURL) GetBranchName() string  { return az.branch }
func (az *AzureURL) GetTag() string         { return az.tag }
func (az *AzureURL) GetOwnerName() string   { return az.owner }
func (az *AzureURL) GetRepoName() string    { return az.repo }
func (az *AzureURL) GetPath() string        { return az.path }
func (az *AzureURL) GetToken() string       { return az.token }
func (az *AzureURL) GetHttpCloneURL() string {
	return fmt.Sprintf("https://%s@%s/%s/%s/_git/%s", az.GetOwnerName(), az.GetHostName(), az.GetOwnerName(), az.GetProjectName(), az.GetRepoName())
}

func (az *AzureURL) SetOwnerName(o string)         { az.owner = o }
func (az *AzureURL) SetProjectName(project string) { az.project = project }
func (az *AzureURL) SetRepoName(r string)          { az.repo = r }
func (az *AzureURL) SetBranchName(branch string)   { az.branch = branch }
func (az *AzureURL) SetTag(tag string)             { az.tag = tag }
func (az *AzureURL) SetPath(p string)              { az.path = p }
func (az *AzureURL) SetToken(token string)         { az.token = token }

// Parse URL
func (az *AzureURL) Parse(fullURL string) error {
	parsedURL, err := giturl.Parse(fullURL)
	if err != nil {
		return err
	}
	az.host = parsedURL.Host

	if strings.HasPrefix(az.host, "ssh") {
		az.host = strings.TrimPrefix(az.host, "ssh.")
		return az.parseHostSSH(parsedURL)
	}
	return az.parseHostHTTP(parsedURL)
}

func (az *AzureURL) parseHostSSH(parsedURL *url.URL) error {
	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice

	if len(splittedRepo) < 3 || len(splittedRepo) > 5 {
		return fmt.Errorf("expecting v/<user>/<project>/<repo> in url path, received: '%s'", parsedURL.Path)
	}

	index := 0
	if len(splittedRepo) == 4 {
		index++
	}
	az.owner = splittedRepo[index]
	az.project = splittedRepo[index+1]
	az.repo = splittedRepo[index+2]

	return nil
}
func (az *AzureURL) parseHostHTTP(parsedURL *url.URL) error {
	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice
	if len(splittedRepo) < 4 || splittedRepo[2] != "_git" {
		return fmt.Errorf("expecting <user>/<project>/_git/<repo> in url path, received: '%s'", parsedURL.Path)
	}
	az.owner = splittedRepo[0]
	az.project = splittedRepo[1]
	az.repo = splittedRepo[3]

	if v := parsedURL.Query().Get("version"); v != "" {
		if strings.HasPrefix(v, "GB") {
			az.branch = strings.TrimPrefix(v, "GB")
		}
		if strings.HasPrefix(v, "GT") {
			az.tag = strings.TrimPrefix(v, "GT")
		}
	}

	if v := parsedURL.Query().Get("path"); v != "" {
		az.path = v
	}

	return nil
}
