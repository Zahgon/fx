package fxtest

import (
	"go.uber.org/fx"
)

type App struct {
	*fx.App

	tb TB
}

func New(tb TB, opts ...fx.Option) *App { _ = "STUB: not implemented"; return nil }

func (app *App) RequireStart() *App { _ = "STUB: not implemented"; return nil }

func (app *App) RequireStop() { _ = "STUB: not implemented"; return }
