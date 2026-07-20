package paramobject

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Config     ClientConfig
	HTTPClient *http.Client

	Logger *zap.Logger `optional:"true"`
}

func New(p Params) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }
