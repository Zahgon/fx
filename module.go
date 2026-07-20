package fx

import (
	"go.uber.org/dig"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/internal/fxreflect"
)

type container interface {
	Invoke(any, ...dig.InvokeOption) error
	Provide(any, ...dig.ProvideOption) error
	Decorate(any, ...dig.DecorateOption) error
}

func Module(name string, opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

type moduleOption struct {
	name     string
	location fxreflect.Frame
	options  []Option
}

func (o moduleOption) String() string { _ = "STUB: not implemented"; return "" }

func (o moduleOption) apply(mod *module) { _ = "STUB: not implemented"; return }

type module struct {
	parent         *module
	name           string
	trace          []string
	scope          scope
	provides       []provide
	invokes        []invoke
	decorators     []decorator
	modules        []*module
	app            *App
	log            fxevent.Logger
	fallbackLogger fxevent.Logger
	logConstructor *provide
}

type scope interface {
	Decorate(f any, opts ...dig.DecorateOption) error
	Invoke(f any, opts ...dig.InvokeOption) error
	Provide(f any, opts ...dig.ProvideOption) error
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	String() string
}

func (m *module) build(app *App, root *dig.Container) { _ = "STUB: not implemented"; return }

func (m *module) provideAll() { _ = "STUB: not implemented"; return }

func (m *module) provide(p provide) { _ = "STUB: not implemented"; return }

func (m *module) supply(p provide) { _ = "STUB: not implemented"; return }

func (m *module) installAllEventLoggers() { _ = "STUB: not implemented"; return }

func (m *module) installEventLogger(buffer *logBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *module) invokeAll() error { _ = "STUB: not implemented"; return nil }

func (m *module) invoke(i invoke) (err error) { _ = "STUB: not implemented"; return nil }

func (m *module) decorateAll() error { _ = "STUB: not implemented"; return nil }

func (m *module) decorate(d decorator) (err error) { _ = "STUB: not implemented"; return nil }

func (m *module) replace(d decorator) error { _ = "STUB: not implemented"; return nil }
