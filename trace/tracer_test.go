package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	tracer.Trace("こんにちは、traceパッケージです")
	if buf.String() != "こんにちは、traceパッケージです\n" {
		t.Errorf("'%s'という誤った文字列が出力されました。", buf.String())
	}
}
