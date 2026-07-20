package fx

import (
	"go.uber.org/dig"
	"go.uber.org/fx/internal/fxreflect"
)

func Provide(constructors ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

type provideOption struct {
	Targets []any
	Stack   fxreflect.Stack
}

func (o provideOption) apply(mod *module) { _ = "STUB: not implemented"; return }

type privateOption struct{}

var Private = privateOption{}

func (o provideOption) String() string { _ = "STUB: not implemented"; return "" }

func runProvide(c container, p provide, opts ...dig.ProvideOption) error {
	_ = "STUB: not implemented"
	return nil
}
