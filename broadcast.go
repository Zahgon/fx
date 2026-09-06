package fx

import (
	"os"
	"sync"
)

type broadcaster struct {
	m sync.Mutex

	last *ShutdownSignal

	done []chan os.Signal

	wait []chan ShutdownSignal
}

func (b *broadcaster) reset() { _ = "STUB: not implemented"; return }

func (b *broadcaster) Done() <-chan os.Signal { _ = "STUB: not implemented"; return nil }

func (b *broadcaster) Wait() <-chan ShutdownSignal { _ = "STUB: not implemented"; return nil }

func (b *broadcaster) Broadcast(signal ShutdownSignal) error { _ = "STUB: not implemented"; return nil }

func (b *broadcaster) broadcast(
	signal ShutdownSignal,
	anchors ...func(ShutdownSignal) (int, int),
) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (b *broadcaster) broadcastDone(signal ShutdownSignal) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (b *broadcaster) broadcastWait(signal ShutdownSignal) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

type unsentSignalError struct {
	Signal ShutdownSignal
	Unsent int
	Total  int
}

func (err *unsentSignalError) Error() string { _ = "STUB: not implemented"; return "" }
