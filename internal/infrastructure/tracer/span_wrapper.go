package tracer

type ISpan interface {
	End()
	AddEvent(name string, opts ...interface{})
}
