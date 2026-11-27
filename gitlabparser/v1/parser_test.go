package v1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitHubParserWithURL(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		url     string
		want    *GitLabURL
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "general case",
			url:  "https://gitlab.com/kubescape/testing",
			want: &GitLabURL{
				host:  "gitlab.com",
				owner: "kubescape",
				repo:  "testing",
			},
			wantErr: assert.NoError,
		},
		{
			name: "file",
			url:  "https://gitlab.com/kubescape/testing/-/blob/main/stable/acs-engine-autoscaler/Chart.yaml",
			want: &GitLabURL{
				host:   "gitlab.com",
				owner:  "kubescape",
				repo:   "testing",
				branch: "main",
				path:   "stable/acs-engine-autoscaler/Chart.yaml",
			},
			wantErr: assert.NoError,
		},
		{
			name: "branch",
			url:  "https://gitlab.com/kubescape/testing/-/blob/dev/README.md",
			want: &GitLabURL{
				host:   "gitlab.com",
				owner:  "kubescape",
				repo:   "testing",
				branch: "dev",
				path:   "README.md",
			},
			wantErr: assert.NoError,
		},
		{
			name: "branch",
			url:  "https://gitlab.com/kubescape/testing/-/tree/dev",
			want: &GitLabURL{
				host:   "gitlab.com",
				owner:  "kubescape",
				repo:   "testing",
				branch: "dev",
			},
			wantErr: assert.NoError,
		},
		{
			name: "tag",
			url:  "https://gitlab.com/kubescape/testing/-/blob/v0.0.0/README.md",
			want: &GitLabURL{
				host:   "gitlab.com",
				owner:  "kubescape",
				repo:   "testing",
				branch: "v0.0.0",
				path:   "README.md",
			},
			wantErr: assert.NoError,
		},
		{
			name: "raw",
			url:  "https://gitlab.com/kubescape/testing/-/raw/main/stable/acs-engine-autoscaler/Chart.yaml",
			want: &GitLabURL{
				host:   "gitlab.com",
				owner:  "kubescape",
				repo:   "testing",
				branch: "main",
				path:   "stable/acs-engine-autoscaler/Chart.yaml",
			},
			wantErr: assert.NoError,
		},
		{
			name: "scp-like syntax",
			url:  "git@gitlab.com:owner/repo.git",
			want: &GitLabURL{
				host:  "gitlab.com",
				owner: "owner",
				repo:  "repo",
			},
			wantErr: assert.NoError,
		},
		{
			name: "project and subproject",
			url:  "https://gitlab.com/matthyx1/subgroup1/project1.git",
			want: &GitLabURL{
				host:  "gitlab.com",
				owner: "matthyx1/subgroup1",
				repo:  "project1",
			},
			wantErr: assert.NoError,
		},
		{
			name: "project and subproject",
			url:  "https://gitlab.com/matthyx1/subgroup1/subsubgroup1/project1.git",
			want: &GitLabURL{
				host:  "gitlab.com",
				owner: "matthyx1/subgroup1/subsubgroup1",
				repo:  "project1",
			},
			wantErr: assert.NoError,
		},
		{
			name: "self-hosted",
			host: "gitlab.host.com",
			url:  "https://gitlab.host.com/kubescape/testing",
			want: &GitLabURL{
				configuredHost: "gitlab.host.com",
				host:           "gitlab.host.com",
				owner:          "kubescape",
				repo:           "testing",
			},
			wantErr: assert.NoError,
		},
		{
			name: "repo only",
			host: "git.host.com",
			url:  "https://git.host.com/repo.git",
			want: &GitLabURL{
				configuredHost: "git.host.com",
				host:           "git.host.com",
				owner:          "",
				repo:           "repo",
			},
			wantErr: assert.NoError,
		},
		{
			name: "repo only with default port",
			host: "git.host.com",
			url:  "https://git.host.com:443/repo.git",
			want: &GitLabURL{
				configuredHost: "git.host.com",
				host:           "git.host.com",
				owner:          "",
				repo:           "repo",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGitLabParserWithURL(tt.host, tt.url)
			if !tt.wantErr(t, err, fmt.Sprintf("NewGitLabParserWithURL(%v)", tt.url)) {
				return
			}
			got.gitLabAPI = nil
			assert.Equalf(t, tt.want, got, "NewGitLabParserWithURL(%v)", tt.url)
		})
	}
}

func TestIsHostGitLab(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "gitlab.com",
			want: true,
		},
		{
			name: "gitlab.host.com",
			want: false,
		},
		{
			name: "github.com",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsHostGitLab(tt.name), "IsHostGitLab(%v)", tt.name)
		})
	}
}
