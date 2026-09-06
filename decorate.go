package fx

import (
	"reflect"

	"go.uber.org/dig"
	"go.uber.org/fx/internal/fxreflect"
)

func Decorate(decorators ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

type decorateOption struct {
	Targets []any
	Stack   fxreflect.Stack
}

func (o decorateOption) apply(mod *module) { _ = "STUB: not implemented"; return }

func (o decorateOption) String() string { _ = "STUB: not implemented"; return "" }

type decorator struct {
	Target any

	Stack fxreflect.Stack

	IsReplace   bool
	ReplaceType reflect.Type
}

func runDecorator(c container, d decorator, opts ...dig.DecorateOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}
