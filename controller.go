// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func (rest *Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (rest *Controller) DELETE(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) GET(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) HEAD(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) OPTIONS(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) PATCH(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) POST(ctx *Context) {
	ctx.ResponseForbidden()
}

func (rest *Controller) PUT(ctx *Context) {
	ctx.ResponseForbidden()
}
