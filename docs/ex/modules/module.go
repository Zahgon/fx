package modules

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("server",

	fx.Provide(
		New,
	),

	fx.Provide(
		fx.Private,
		parseConfig,
	),

	fx.Invoke(startServer),

	fx.Decorate(wrapLogger),
)

type Config struct {
	Addr string `yaml:"addr"`
}

func parseConfig() (Config, error) { _ = "STUB: not implemented"; return *new(Config), nil }

type Params struct {
	fx.In

	Log    *zap.Logger
	Config Config
}

type Result struct {
	fx.Out

	Server *Server
}

func New(p Params) (Result, error) { _ = "STUB: not implemented"; return *new(Result), nil }

type Server struct{}

func (*Server) Start() error { _ = "STUB: not implemented"; return nil }

func startServer(srv *Server) error { _ = "STUB: not implemented"; return nil }

func wrapLogger(log *zap.Logger) *zap.Logger { _ = "STUB: not implemented"; return nil }
