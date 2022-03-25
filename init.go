package urlgit

import (
	"fmt"
	"net/url"
	githubparserv1 "urlgit/githubparser/v1"
)

// NewGitURL get instance of git parser
func NewGitURL(fullURL string) (IGitURL, error) {
	hostUrl, err := getHost(fullURL)
	if err != nil {
		return nil, err
	}

	switch hostUrl {
	case "github.com", "raw.githubusercontent.com":
		return githubparserv1.NewGitHubParserWithURL(fullURL)
	default:
		return nil, fmt.Errorf("repository host '%s' not supported", hostUrl)
	}
}

func getHost(fullURL string) (string, error) {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host, nil
}
