package fx

import "io"

type printerWriter struct{ p Printer }

func writerFromPrinter(p Printer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (w *printerWriter) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
