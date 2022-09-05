package apis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	urlAzure  = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public"
	urlGitHub = "https://github.com/kubescape/go-git-url"
)

func TestNewUrlComposer(t *testing.T) {
	{
		var provider = ProviderType(ProviderAzure.String())
		g, err := NewUrlComposer(provider, urlAzure, "rules-tests/k8s-audit-logs-enabled-native/test-failed/input/apiserverpod.json")
		assert.NoError(t, err)
		assert.Equal(t,
			"https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=rules-tests/k8s-audit-logs-enabled-native/test-failed/input/apiserverpod.json&version=GBmaster",
			g.FileUrlByBranch("master"))
		assert.Equal(t,
			"https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=rules-tests/k8s-audit-logs-enabled-native/test-failed/input/apiserverpod.json&version=GC9300d038ca1be76711aacb2f7ba59b67f4bbddf1",
			g.FileUrlByCommit("9300d038ca1be76711aacb2f7ba59b67f4bbddf1"))
		assert.Equal(t,
			"https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=rules-tests/k8s-audit-logs-enabled-native/test-failed/input/apiserverpod.json&version=GTv1.0.179",
			g.FileUrlByTag("v1.0.179"))
	}
	{
		var provider = ProviderType(ProviderGitHub.String())
		g, err := NewUrlComposer(provider, urlGitHub, "files/file0.json")
		assert.NoError(t, err)
		assert.Equal(t,
			"https://github.com/kubescape/go-git-url/blob/master/files/file0.json",
			g.FileUrlByBranch("master"))
		assert.Equal(t,
			"https://github.com/kubescape/go-git-url/blob/c7380d9b941a671253a81c694f644868270a1d81/files/file0.json",
			g.FileUrlByCommit("c7380d9b941a671253a81c694f644868270a1d81"))
		assert.Equal(t,
			"https://github.com/kubescape/go-git-url/blob/v0.0.13/files/file0.json",
			g.FileUrlByTag("v0.0.13"))
	}

}
