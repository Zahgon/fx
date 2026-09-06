package fx

import (
	"context"

	"go.uber.org/fx/internal/lifecycle"
)

type HookFunc interface {
	~func() | ~func() error | ~func(context.Context) | ~func(context.Context) error
}

type Lifecycle interface {
	Append(Hook)
}

type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error

	onStartName string
	onStopName  string
}

func StartHook[T HookFunc](start T) Hook { _ = "STUB: not implemented"; return *new(Hook) }

func StopHook[T HookFunc](stop T) Hook { _ = "STUB: not implemented"; return *new(Hook) }

func StartStopHook[T1 HookFunc, T2 HookFunc](start T1, stop T2) Hook {
	_ = "STUB: not implemented"
	return *new(Hook)
}

type lifecycleWrapper struct {
	*lifecycle.Lifecycle
}

func (l *lifecycleWrapper) Append(h Hook) { _ = "STUB: not implemented"; return }
