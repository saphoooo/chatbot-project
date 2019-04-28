package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describe an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// New blablabla
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
