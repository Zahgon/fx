//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package fx

import "golang.org/x/sys/unix"

const (
	_sigINT  = unix.SIGINT
	_sigTERM = unix.SIGTERM
)
