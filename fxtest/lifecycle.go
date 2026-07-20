package fxtest

import (
	"context"
	"io"

	"go.uber.org/fx"
	"go.uber.org/fx/internal/lifecycle"
)

type panicT struct {
	W io.Writer

	lastErr string
}

var _ TB = &panicT{}

func (t *panicT) format(s string, args ...any) string { _ = "STUB: not implemented"; return "" }

func (t *panicT) Logf(s string, args ...any) { _ = "STUB: not implemented"; return }

func (t *panicT) Errorf(s string, args ...any) { _ = "STUB: not implemented"; return }

func (t *panicT) FailNow() { _ = "STUB: not implemented"; return }

type LifecycleOption interface {
	apply(*Lifecycle)
}

func EnforceTimeout(enforce bool) LifecycleOption {
	_ = "STUB: not implemented"
	return *new(LifecycleOption)
}

type enforceTimeout struct {
	enforce bool
}

func (e *enforceTimeout) apply(lc *Lifecycle) { _ = "STUB: not implemented"; return }

var _ LifecycleOption = (*enforceTimeout)(nil)

type Lifecycle struct {
	t  TB
	lc *lifecycle.Lifecycle

	enforceTimeout bool
}

var _ fx.Lifecycle = (*Lifecycle)(nil)

func NewLifecycle(t TB, opts ...LifecycleOption) *Lifecycle { _ = "STUB: not implemented"; return nil }

func (l *Lifecycle) withTimeout(ctx context.Context, fn func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Lifecycle) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *Lifecycle) RequireStart() *Lifecycle { _ = "STUB: not implemented"; return nil }

func (l *Lifecycle) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *Lifecycle) RequireStop() { _ = "STUB: not implemented"; return }

func (l *Lifecycle) Append(h fx.Hook) { _ = "STUB: not implemented"; return }
