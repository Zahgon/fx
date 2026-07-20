package paramobject

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Client struct {
	url  string
	http *http.Client
	log  *zap.Logger
}

type ClientConfig struct {
	URL string
}

type ClientParams struct {
	fx.In

	Config     ClientConfig
	HTTPClient *http.Client
}

func NewClient(p ClientParams) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }
