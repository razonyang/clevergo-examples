# RESTFul API Controller
RESTFul API Controller, the controller is a middleware.

### Controller Interface
```
type ControllerInterface interface {
	Handle(next Handler) Handler

	DELETE(ctx *Context)
	GET(ctx *Context)
	HEAD(ctx *Context)
	OPTIONS(ctx *Context)
	PATCH(ctx *Context)
	POST(ctx *Context)
	PUT(ctx *Context)
}
```

### Create a RESTFul API Controller
```
type MyController struct {
    clevergo.Controller
}

func (c MyController) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c MyController) GET(ctx *Context) {
	ctx.SetBodyString("RESTFul API Controller.")
}
```
The other request like POST, DELETE etc, will response `forbidden` to client.

### Register Controller
`Router.RegisterController(MyController{})`

### Shortcut
- [Router](router.md)
- [Context](context.md)
- [Handler](handler.md)
- [Middleware](middleware.md)
- [Examples](/examples)