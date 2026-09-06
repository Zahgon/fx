package feed

import "go.uber.org/fx"

var ResultModule = fx.Options(

	fx.Provide(New),
)

type Watcher interface{}

type watcher struct{}

type Result struct {
	fx.Out

	Watcher Watcher `group:"watchers"`
}

func New() (Result, error) { _ = "STUB: not implemented"; return *new(Result), nil }
