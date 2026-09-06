package feed

import "go.uber.org/fx"

var AnnotateModule = fx.Options(

	fx.Provide(
		NewWatcher,
	),

	fx.Provide(

		fx.Annotate(
			NewWatcher,

			fx.ResultTags(`group:"watchers"`),
		),
	),
)

type FileWatcher struct{}

var FileWatcherModule = fx.Options(
	fx.Provide(

		fx.Annotate(
			NewFileWatcher,
			fx.As(new(Watcher)),
			fx.ResultTags(`group:"watchers"`),
		),
	),
)

func NewFileWatcher() (*FileWatcher, error) { _ = "STUB: not implemented"; return nil, nil }

func NewWatcher() (Watcher, error) { _ = "STUB: not implemented"; return *new(Watcher), nil }
