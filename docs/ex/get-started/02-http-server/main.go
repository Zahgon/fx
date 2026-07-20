package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(NewHTTPServer),

		fx.Invoke(func(*http.Server) {}),
	).Run()

}

func NewHTTPServer(lc fx.Lifecycle) *http.Server { _ = "STUB: not implemented"; return nil }
