package fxtest

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewTestLogger(t TB) fxevent.Logger { _ = "STUB: not implemented"; return *new(fxevent.Logger) }

func WithTestLogger(t TB) fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

type testPrinter struct {
	TB
}

func NewTestPrinter(t TB) fx.Printer { _ = "STUB: not implemented"; return *new(fx.Printer) }

func (p *testPrinter) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }
