package fx

import (
	"go.uber.org/fx/internal/fxreflect"
)

func Invoke(funcs ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

type invokeOption struct {
	Targets []any
	Stack   fxreflect.Stack
}

func (o invokeOption) apply(mod *module) { _ = "STUB: not implemented"; return }

func (o invokeOption) String() string { _ = "STUB: not implemented"; return "" }

func runInvoke(c container, i invoke) error { _ = "STUB: not implemented"; return nil }
