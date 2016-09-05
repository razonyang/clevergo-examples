package clevergo

type Handler interface {
	Handle(*Context)
}

type HandlerFunc func(*Context)

func (h HandlerFunc) Handle(ctx *Context) {
	h(ctx)
}
