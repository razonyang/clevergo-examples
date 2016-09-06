package clevergo

// A Handler responds to an HTTP request.
type Handler interface {
	Handle(*Context)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.
type HandlerFunc func(*Context)

// Handle calls f(ctx).
func (f HandlerFunc) Handle(ctx *Context) {
	f(ctx)
}
