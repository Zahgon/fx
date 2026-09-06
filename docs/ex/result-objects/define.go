package resultobject

import "go.uber.org/fx"

type Client struct{}

type ClientResult struct {
	fx.Out

	Client *Client
}

func NewClient() (ClientResult, error) { _ = "STUB: not implemented"; return *new(ClientResult), nil }
