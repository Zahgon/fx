package consume

import "go.uber.org/fx"

var PlainModule = fx.Options(

	fx.Provide(
		NewEmitter,
	),
)

var AnnotateModule = fx.Options(

	fx.Provide(

		fx.Annotate(
			NewEmitter,

			fx.ParamTags(`group:"watchers"`),
		),
	),
)

type Emitter struct{ ws []Watcher }

func NewEmitter(watchers []Watcher) (*Emitter, error) { _ = "STUB: not implemented"; return nil, nil }

var EmitterFromModule = fx.Options(
	fx.Provide(

		fx.Annotate(
			EmitterFrom,
			fx.ParamTags(`group:"watchers"`),
		),
	),
)

func EmitterFrom(watchers ...Watcher) (*Emitter, error) { _ = "STUB: not implemented"; return nil, nil }
