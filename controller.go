package clevergo

// ControllerInterface.
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

// Controller is an empty struct.
type Controller struct{}

// Handle implemented Middleware Interface.
func (c Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

// DELETE for handling the DELETE request.
func (c Controller) DELETE(ctx *Context) {
	ctx.NotFound()
}

// GET for handling the GET request.
func (c Controller) GET(ctx *Context) {
	ctx.NotFound()
}

// HEAD for handling the HEAD request.
func (c Controller) HEAD(ctx *Context) {
	ctx.NotFound()
}

// OPTIONS for handling the OPTIONS request.
func (c Controller) OPTIONS(ctx *Context) {
	ctx.NotFound()
}

// PATCH for handling the PATCH request.
func (c Controller) PATCH(ctx *Context) {
	ctx.NotFound()
}

// POST for handling the POST request.
func (c Controller) POST(ctx *Context) {
	ctx.NotFound()
}

// PUT for handling the PUT request.
func (c Controller) PUT(ctx *Context) {
	ctx.NotFound()
}
