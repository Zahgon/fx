package fx

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.uber.org/dig"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/internal/fxclock"
	"go.uber.org/fx/internal/fxreflect"
)

const DefaultTimeout = 15 * time.Second

type Option interface {
	fmt.Stringer

	apply(*module)
}

func Error(errs ...error) Option { _ = "STUB: not implemented"; return *new(Option) }

type errorOption []error

func (errs errorOption) apply(mod *module) { _ = "STUB: not implemented"; return }

func (errs errorOption) String() string { _ = "STUB: not implemented"; return "" }

func Options(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

type optionGroup []Option

func (og optionGroup) apply(mod *module) { _ = "STUB: not implemented"; return }

func (og optionGroup) String() string { _ = "STUB: not implemented"; return "" }

func StartTimeout(v time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type startTimeoutOption time.Duration

func (t startTimeoutOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (t startTimeoutOption) String() string { _ = "STUB: not implemented"; return "" }

func StopTimeout(v time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type stopTimeoutOption time.Duration

func (t stopTimeoutOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (t stopTimeoutOption) String() string { _ = "STUB: not implemented"; return "" }

func RecoverFromPanics() Option { _ = "STUB: not implemented"; return *new(Option) }

type recoverFromPanicsOption struct{}

func (o recoverFromPanicsOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (o recoverFromPanicsOption) String() string { _ = "STUB: not implemented"; return "" }

func WithLogger(constructor any) Option { _ = "STUB: not implemented"; return *new(Option) }

type withLoggerOption struct {
	constructor any
	Stack       fxreflect.Stack
}

func (l withLoggerOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (l withLoggerOption) String() string { _ = "STUB: not implemented"; return "" }

type Printer interface {
	Printf(string, ...any)
}

func Logger(p Printer) Option { _ = "STUB: not implemented"; return *new(Option) }

type loggerOption struct{ p Printer }

func (l loggerOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (l loggerOption) String() string { _ = "STUB: not implemented"; return "" }

var NopLogger = WithLogger(func() fxevent.Logger { return fxevent.NopLogger })

type App struct {
	err       error
	clock     fxclock.Clock
	lifecycle *lifecycleWrapper

	container *dig.Container
	root      *module

	startTimeout time.Duration
	stopTimeout  time.Duration

	errorHooks []ErrorHandler
	validate   bool

	recoverFromPanics bool

	receivers signalReceivers

	osExit func(code int)
}

type provide struct {
	Target any

	Stack fxreflect.Stack

	IsSupply   bool
	SupplyType reflect.Type

	Private bool
}

type invoke struct {
	Target any

	Stack fxreflect.Stack
}

type ErrorHandler interface {
	HandleError(error)
}

func ErrorHook(funcs ...ErrorHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

type errorHookOption []ErrorHandler

func (eho errorHookOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (eho errorHookOption) String() string { _ = "STUB: not implemented"; return "" }

type errorHandlerList []ErrorHandler

func (ehl errorHandlerList) HandleError(err error) { _ = "STUB: not implemented"; return }

func validate(validate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type validateOption struct {
	validate bool
}

func (o validateOption) apply(m *module) { _ = "STUB: not implemented"; return }

func (o validateOption) String() string { _ = "STUB: not implemented"; return "" }

func ValidateApp(opts ...Option) error { _ = "STUB: not implemented"; return nil }

func New(opts ...Option) *App { _ = "STUB: not implemented"; return nil }

func (app *App) log() fxevent.Logger { _ = "STUB: not implemented"; return *new(fxevent.Logger) }

type DotGraph string

type errWithGraph interface {
	Graph() DotGraph
}

type errorWithGraph struct {
	graph string
	err   error
}

func (err errorWithGraph) Graph() DotGraph { _ = "STUB: not implemented"; return *new(DotGraph) }

func (err errorWithGraph) Error() string { _ = "STUB: not implemented"; return "" }

func VisualizeError(err error) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (app *App) exit(code int) { _ = "STUB: not implemented"; return }

func (app *App) Run() { _ = "STUB: not implemented"; return }

func (app *App) run(done func() <-chan ShutdownSignal) (exitCode int) {
	_ = "STUB: not implemented"
	return 0
}

func (app *App) Err() error { _ = "STUB: not implemented"; return nil }

var (
	_onStartHook = "OnStart"
	_onStopHook  = "OnStop"
)

func (app *App) Start(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (app *App) withRollback(
	ctx context.Context,
	f func(context.Context) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (app *App) Stop(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (app *App) Done() <-chan os.Signal { _ = "STUB: not implemented"; return nil }

func (app *App) Wait() <-chan ShutdownSignal { _ = "STUB: not implemented"; return nil }

func (app *App) StartTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (app *App) StopTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (app *App) dotGraph() (DotGraph, error) { _ = "STUB: not implemented"; return *new(DotGraph), nil }

type withTimeoutParams struct {
	log       fxevent.Logger
	hook      string
	callback  func(context.Context) error
	lifecycle *lifecycleWrapper
}

var errHookCallbackExited = errors.New("goroutine exited without returning")

func withTimeout(ctx context.Context, param *withTimeoutParams) error {
	_ = "STUB: not implemented"
	return nil
}

type appLogger struct{ app *App }

func (l appLogger) LogEvent(ev fxevent.Event) { _ = "STUB: not implemented"; return }
