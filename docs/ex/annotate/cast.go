package annotate

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/docs/ex/annotate/github"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

type Config struct{}

func NewHTTPClient(Config) (*http.Client, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGitHubClient(client HTTPClient) *github.Client { _ = "STUB: not implemented"; return nil }

func options() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }
