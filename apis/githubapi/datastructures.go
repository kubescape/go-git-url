package githubapi

type innerTree struct {
	Path string `json:"path"`
}
type Tree struct {
	InnerTrees []innerTree `json:"tree"`
}

type githubDefaultBranchAPI struct {
	DefaultBranch string `json:"default_branch"`
}

type Headres struct {
	Token string
}
