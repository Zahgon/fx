//go:build windows
// +build windows

package fx

import "golang.org/x/sys/windows"

const (
	_sigINT  = windows.SIGINT
	_sigTERM = windows.SIGTERM
)
