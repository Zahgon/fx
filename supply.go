package fx

import (
	"reflect"

	"go.uber.org/fx/internal/fxreflect"
)

func Supply(values ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

type supplyOption struct {
	Targets []any
	Types   []reflect.Type
	Stack   fxreflect.Stack
	Private bool
}

func (o supplyOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (o supplyOption) String() string { _ = "STUB: not implemented"; return "" }

func newSupplyConstructor(value any) (any, reflect.Type) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Type)
}
