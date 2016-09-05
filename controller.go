package clevergo

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

type Controller struct {
}

func (c Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c Controller) DELETE(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) GET(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) HEAD(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) PATCH(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) POST(ctx *Context) {
	ctx.NotFound()
}

func (c Controller) PUT(ctx *Context) {
	ctx.NotFound()
}

type RESTController struct {
	AllowOrigin  string
	AllowMethods string
}

func (c RESTController) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Set Access-Control-Allow-Origin and Access-Control-Allow-Methods for ajax request.
		ctx.Response.Header.Set("Access-Control-Allow-Origin", c.AllowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", c.AllowMethods)

		// Using param named '_method' to simulate the other request, such as PUT, DELETE etc.
		if !ctx.IsGet() {
			switch string(ctx.FormValue("_method")) {
			case "PUT":
				c.PUT(ctx)
				return
			case "DELETE":
				c.DELETE(ctx)
				return
			case "HEAD":
				c.HEAD(ctx)
				return
			case "OPTIONS":
				c.OPTIONS(ctx)
				return
			case "PATCH":
				c.PATCH(ctx)
				return
			}
		}

		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c RESTController) DELETE(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) GET(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) HEAD(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) PATCH(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) POST(ctx *Context) {
	ctx.NotFound()
}

func (c RESTController) PUT(ctx *Context) {
	ctx.NotFound()
}
