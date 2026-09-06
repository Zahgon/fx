package consume

import "go.uber.org/fx"

type Watcher interface{}

var ParamsModule = fx.Options(

	fx.Provide(New),
)

type Params struct {
	fx.In

	Watchers []Watcher `group:"watchers"`
}

type Result struct {
	fx.Out

	Emitter *Emitter
}

func New(p Params) (Result, error) { _ = "STUB: not implemented"; return *new(Result), nil }
