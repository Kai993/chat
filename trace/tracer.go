package trace

import (
	"fmt"
	"io"
)

// Tracer : コード内での出来事を記録できるオブジェクトを表すインタフェース
type Tracer interface {
	Trace(...interface{})
}

// New : Tracerを生成する
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off : ログを無効化する
func Off() Tracer {
	return &nilTracer{}
}
