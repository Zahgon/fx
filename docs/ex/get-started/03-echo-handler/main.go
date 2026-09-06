package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(

		fx.Provide(
			NewHTTPServer,

			NewServeMux,

			NewEchoHandler,
		),

		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewServeMux(echo *EchoHandler) *http.ServeMux { _ = "STUB: not implemented"; return nil }

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler { _ = "STUB: not implemented"; return nil }

func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	_ = "STUB: not implemented"
	return nil
}
