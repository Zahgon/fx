package fx

import (
	"reflect"

	"go.uber.org/fx/internal/fxreflect"
)

func Replace(values ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

type replaceOption struct {
	Targets []any
	Types   []reflect.Type
	Stack   fxreflect.Stack
}

func (o replaceOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (o replaceOption) String() string { _ = "STUB: not implemented"; return "" }

func newReplaceDecorator(value any) (any, reflect.Type) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Type)
}
