package clevergo

// Controller Interface.
//
// In fact, the controller is a middleware.
type ControllerInterface interface {
	Handle(next Handler) Handler // Implemented Middleware Interface.

	DELETE(ctx *Context)  // Request handler for DELETE request.
	GET(ctx *Context)     // Request handler for GET request.
	HEAD(ctx *Context)    // Request handler for HEAD request.
	OPTIONS(ctx *Context) // Request handler for OPTIONS request.
	PATCH(ctx *Context)   // Request handler for PATCH request.
	POST(ctx *Context)    // Request handler for POST request.
	PUT(ctx *Context)     // Request handler for PUT request.
}

// Controller.
type Controller struct{}

// Controller handler.
//
// Implemented Middleware Interface.
func (c Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

// Request handler for DELETE request.
func (c Controller) DELETE(ctx *Context) {
	ctx.NotFound()
}

// Request handler for GET request.
func (c Controller) GET(ctx *Context) {
	ctx.NotFound()
}

// Request handler for HEAD request.
func (c Controller) HEAD(ctx *Context) {
	ctx.NotFound()
}

// Request handler for OPTIONS request.
func (c Controller) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

// Request handler for PATCH request.
func (c Controller) PATCH(ctx *Context) {
	ctx.NotFound()
}

// Request handler for POST request.
func (c Controller) POST(ctx *Context) {
	ctx.NotFound()
}

// Request handler for PUT request.
func (c Controller) PUT(ctx *Context) {
	ctx.NotFound()
}

// RESTful API Controller.
type RESTController struct {
	AllowOrigin  string
	AllowMethods string
}

// RESTful API Controller handler.
//
// Implemented Middleware Interface.
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

// Request handler for DELETE request.
func (c RESTController) DELETE(ctx *Context) {
	ctx.NotFound()
}

// Request handler for GET request.
func (c RESTController) GET(ctx *Context) {
	ctx.NotFound()
}

// Request handler for HEAD request.
func (c RESTController) HEAD(ctx *Context) {
	ctx.NotFound()
}

// Request handler for OPTIONS request.
func (c RESTController) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

// Request handler for PATCH request.
func (c RESTController) PATCH(ctx *Context) {
	ctx.NotFound()
}

// Request handler for POST request.
func (c RESTController) POST(ctx *Context) {
	ctx.NotFound()
}

// Request handler for PUT request.
func (c RESTController) PUT(ctx *Context) {
	ctx.NotFound()
}
