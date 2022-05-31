package githubapi

import "time"

type ObjectType string

const (
	ObjectTypeDir  ObjectType = "tree"
	ObjectTypeFile ObjectType = "blob"
)

type InnerTree struct {
	Path string     `json:"path"`
	Mode string     `json:"mode"`
	SHA  string     `json:"sha"`
	URL  string     `json:"url"`
	Type ObjectType `json:"type"`
}
type Tree struct {
	InnerTrees []InnerTree `json:"tree"`
	SHA        string      `json:"sha"`
	URL        string      `json:"url"`
	Truncated  bool        `json:"truncated"`
}

type githubDefaultBranchAPI struct {
	DefaultBranch string `json:"default_branch"`
}

type Headers struct {
	Token string
}

// LatestCommit returned structure
type Commit struct {
	SHA         string          `json:"sha"`
	NodeID      string          `json:"node_id"`
	Commit      CommitsMetadata `json:"commit"`
	URL         string          `json:"url"`
	HtmlURL     string          `json:"html_url"`
	CommentsURL string          `json:"comments_url"`
	Author      Author          `json:"author"`
	Committer   Committer       `json:"committer"`
	Parents     []struct {
		SHA     string `json:"sha"`
		URL     string `json:"url"`
		HtmlURL string `json:"html_url"`
	} `json:"parents"`
	Stats struct {
		Total     int `json:"total"`
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
	} `json:"stats"`
	Files []Files `json:"files"`
}

type Files struct {
	SHA         string `json:"sha"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	Changes     int    `json:"changes"`
	BlobURL     string `json:"blob_url"`
	RawURL      string `json:"raw_url"`
	ContentsURL string `json:"contents_url"`
	Patch       string `json:"patch"`
}
type CommitsMetadata struct {
	Author struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"author"`
	Committer struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"committer"`
	Message string `json:"message"`
	Tree    struct {
		SHA string `json:"sha"`
		URL string `json:"url"`
	} `json:"tree"`
	URL          string `json:"url"`
	CommentCount int    `json:"comment_count"`
	Verification struct {
		Verified  bool        `json:"verified"`
		Reason    string      `json:"reason"`
		Signature interface{} `json:"signature"`
		Payload   interface{} `json:"payload"`
	} `json:"verification"`
}
type Author struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Committer struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
