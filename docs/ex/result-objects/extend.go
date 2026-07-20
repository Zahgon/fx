package resultobject

import "go.uber.org/fx"

type Inspector struct{}

type Result struct {
	fx.Out

	Client *Client

	Inspector *Inspector
}

func New() (Result, error) { _ = "STUB: not implemented"; return *new(Result), nil }
