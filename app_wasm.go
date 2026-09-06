//go:build (js && wasm) || (wasip1 && wasm)
// +build js,wasm wasip1,wasm

package fx

import "syscall"

const (
	_sigINT  = syscall.SIGINT
	_sigTERM = syscall.SIGTERM
)
