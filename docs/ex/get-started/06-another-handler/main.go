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
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`name:"echo"`, `name:"hello"`),
			),

			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),

				fx.ResultTags(`name:"echo"`),
			),
			fx.Annotate(
				NewHelloHandler,
				fx.As(new(Route)),

				fx.ResultTags(`name:"hello"`),
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

func NewServeMux(route1, route2 Route) *http.ServeMux { _ = "STUB: not implemented"; return nil }

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler { _ = "STUB: not implemented"; return nil }

func (*HelloHandler) Pattern() string { _ = "STUB: not implemented"; return "" }

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

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
