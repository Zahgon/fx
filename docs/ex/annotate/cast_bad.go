//go:build ignore
// +build ignore

package annotate

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/docs/ex/annotate/github"
)

func NewGitHubClient(client *http.Client) *github.Client { _ = "STUB: not implemented"; return nil }

func options() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }
