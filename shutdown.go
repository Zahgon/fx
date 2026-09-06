package fx

import (
	"time"
)

type Shutdowner interface {
	Shutdown(...ShutdownOption) error
}

type ShutdownOption interface {
	apply(*shutdowner)
}

type exitCodeOption int

func (code exitCodeOption) apply(s *shutdowner) { _ = "STUB: not implemented"; return }

var _ ShutdownOption = exitCodeOption(0)

func ExitCode(code int) ShutdownOption { _ = "STUB: not implemented"; return *new(ShutdownOption) }

type shutdownTimeoutOption time.Duration

func (shutdownTimeoutOption) apply(*shutdowner) { _ = "STUB: not implemented"; return }

var _ ShutdownOption = shutdownTimeoutOption(0)

func ShutdownTimeout(timeout time.Duration) ShutdownOption {
	_ = "STUB: not implemented"
	return *new(ShutdownOption)
}

type shutdowner struct {
	app      *App
	exitCode int
}

func (s *shutdowner) Shutdown(opts ...ShutdownOption) error { _ = "STUB: not implemented"; return nil }

func (app *App) shutdowner() Shutdowner { _ = "STUB: not implemented"; return *new(Shutdowner) }
