package fxtest

type TB interface {
	Logf(string, ...any)
	Errorf(string, ...any)
	FailNow()
}
