package fx

import (
	"context"
	"os"
	"sync"
)

type ShutdownSignal struct {
	Signal   os.Signal
	ExitCode int
}

func (sig ShutdownSignal) String() string { _ = "STUB: not implemented"; return "" }

func newSignalReceivers() signalReceivers { _ = "STUB: not implemented"; return *new(signalReceivers) }

type signalReceivers struct {
	m sync.Mutex

	signals chan os.Signal

	shutdown chan struct{}

	finished chan struct{}

	notify     func(c chan<- os.Signal, sig ...os.Signal)
	stopNotify func(c chan<- os.Signal)

	b *broadcaster
}

func (recv *signalReceivers) relayer() { _ = "STUB: not implemented"; return }

func (recv *signalReceivers) running() bool { _ = "STUB: not implemented"; return false }

func (recv *signalReceivers) Start() { _ = "STUB: not implemented"; return }

func (recv *signalReceivers) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (recv *signalReceivers) Done() <-chan os.Signal { _ = "STUB: not implemented"; return nil }

func (recv *signalReceivers) Wait() <-chan ShutdownSignal { _ = "STUB: not implemented"; return nil }
