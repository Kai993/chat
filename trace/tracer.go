package trace

import (
	"fmt"
	"io"
)

// Tracer : コード内での出来事を記録する
type Tracer struct {
	out io.Writer
}

// New : Tracerを生成する
func New(w io.Writer) Tracer {
	return Tracer{out: w}
}

// Trace : ログを記録する
func (t Tracer) Trace(a ...interface{}) {
	if t.out == nil {
		return
	}

	fmt.Fprintln(t.out, a...)
}
