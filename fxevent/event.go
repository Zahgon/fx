package fxevent

import (
	"os"
	"time"
)

type Event interface {
	event()
}

func (*OnStartExecuting) event()  { _ = "STUB: not implemented"; return }
func (*OnStartExecuted) event()   { _ = "STUB: not implemented"; return }
func (*OnStopExecuting) event()   { _ = "STUB: not implemented"; return }
func (*OnStopExecuted) event()    { _ = "STUB: not implemented"; return }
func (*Supplied) event()          { _ = "STUB: not implemented"; return }
func (*Provided) event()          { _ = "STUB: not implemented"; return }
func (*Replaced) event()          { _ = "STUB: not implemented"; return }
func (*Decorated) event()         { _ = "STUB: not implemented"; return }
func (*BeforeRun) event()         { _ = "STUB: not implemented"; return }
func (*Run) event()               { _ = "STUB: not implemented"; return }
func (*Invoking) event()          { _ = "STUB: not implemented"; return }
func (*Invoked) event()           { _ = "STUB: not implemented"; return }
func (*Stopping) event()          { _ = "STUB: not implemented"; return }
func (*Stopped) event()           { _ = "STUB: not implemented"; return }
func (*RollingBack) event()       { _ = "STUB: not implemented"; return }
func (*RolledBack) event()        { _ = "STUB: not implemented"; return }
func (*Started) event()           { _ = "STUB: not implemented"; return }
func (*LoggerInitialized) event() { _ = "STUB: not implemented"; return }

type OnStartExecuting struct {
	FunctionName string

	CallerName string
}

type OnStartExecuted struct {
	FunctionName string

	CallerName string

	Method string

	Runtime time.Duration

	Err error
}

type OnStopExecuting struct {
	FunctionName string

	CallerName string
}

type OnStopExecuted struct {
	FunctionName string

	CallerName string

	Runtime time.Duration

	Err error
}

type Supplied struct {
	TypeName string

	StackTrace []string

	ModuleTrace []string

	ModuleName string

	Err error
}

type Provided struct {
	ConstructorName string

	StackTrace []string

	ModuleTrace []string

	OutputTypeNames []string

	ModuleName string

	Err error

	Private bool
}

type Replaced struct {
	OutputTypeNames []string

	StackTrace []string

	ModuleTrace []string

	ModuleName string

	Err error
}

type Decorated struct {
	DecoratorName string

	StackTrace []string

	ModuleTrace []string

	ModuleName string

	OutputTypeNames []string

	Err error
}

type BeforeRun struct {
	Name string

	Kind string

	ModuleName string
}

type Run struct {
	Name string

	Kind string

	ModuleName string

	Runtime time.Duration

	Err error
}

type Invoking struct {
	FunctionName string

	ModuleName string
}

type Invoked struct {
	FunctionName string

	ModuleName string

	Err error

	Trace string
}

type Started struct {
	Err error
}

type Stopping struct {
	Signal os.Signal
}

type Stopped struct {
	Err error
}

type RollingBack struct {
	StartErr error
}

type RolledBack struct {
	Err error
}

type LoggerInitialized struct {
	ConstructorName string

	Err error
}
