package main

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
			),
			zap.NewExample,
		),

		fx.Invoke(func(*http.Server) {}),
	).Run()
}

type Route interface {
	http.Handler

	Pattern() string
}

func NewServeMux(route Route) *http.ServeMux { _ = "STUB: not implemented"; return nil }

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler { _ = "STUB: not implemented"; return nil }

func (*EchoHandler) Pattern() string { _ = "STUB: not implemented"; return "" }

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	_ = "STUB: not implemented"
	return nil
}
