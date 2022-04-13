package githubapi

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

type Headres struct {
	Token string
}
