// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"github.com/clevergo/router"
	"github.com/clevergo/sessions"
	"github.com/valyala/fasthttp"
)

type Router struct {
	*router.Router
	middlewares  []Middleware
	sessionStore sessions.Store
}

func NewRouter() *Router {
	return &Router{
		Router:      router.New(),
		middlewares: make([]Middleware, 0),
	}
}

func (r *Router) SetSessionStore(store sessions.Store) {
	r.sessionStore = store
}

func (r *Router) SetMiddlewares(middlewares []Middleware) {
	r.middlewares = middlewares
}

func (r *Router) AddMiddleware(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) GET(path string, handler Handler) {
	r.Router.GET(path, r.getHandler(handler))
}

func (r *Router) HEAD(path string, handler Handler) {
	r.Router.HEAD(path, r.getHandler(handler))
}

func (r *Router) OPTIONS(path string, handler Handler) {
	r.Router.OPTIONS(path, r.getHandler(handler))
}

func (r *Router) POST(path string, handler Handler) {
	r.Router.POST(path, r.getHandler(handler))
}

func (r *Router) PUT(path string, handler Handler) {
	r.Router.PUT(path, r.getHandler(handler))
}

func (r *Router) PATCH(path string, handler Handler) {
	r.Router.PATCH(path, r.getHandler(handler))
}

func (r *Router) DELETE(path string, handler Handler) {
	r.Router.DELETE(path, r.getHandler(handler))
}

func (r *Router) Handle(method, path string, handler Handler) {
	r.Router.Handle(method, path, r.getHandler(handler))
}

func (r *Router) getHandler(handler Handler) router.Handle {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i].Handle(handler)
	}

	return func(_ctx *fasthttp.RequestCtx, ps router.Params) {
		ctx := NewContext(r, _ctx, &ps)
		defer ctx.Close()
		handler.Handle(ctx)
	}
}

func (r *Router) RegisterController(route string, c ControllerInterface) {
	handlers := make(map[string]Handler, 0)

	// Register GET request's handler.
	var getHandler Handler
	getHandler = c.Handle(HandlerFunc(c.GET))
	handlers["GET"] = getHandler

	// Register POST request's handler.
	var postHandler Handler
	postHandler = c.Handle(HandlerFunc(c.POST))
	handlers["POST"] = postHandler

	// Register DELETE request's handler.
	var deleteHandler Handler
	deleteHandler = c.Handle(HandlerFunc(c.DELETE))
	handlers["DELETE"] = deleteHandler

	// Register PUT request's handler.
	var putHandler Handler
	putHandler = c.Handle(HandlerFunc(c.PUT))
	handlers["PUT"] = putHandler

	// Register OPTIONS request's handler.
	var optionsHandler Handler
	optionsHandler = c.Handle(HandlerFunc(c.OPTIONS))
	handlers["OPTIONS"] = optionsHandler

	// Register PATCH request's handler.
	var patchHandler Handler
	patchHandler = c.Handle(HandlerFunc(c.PATCH))
	handlers["PATCH"] = patchHandler

	for method, handler := range handlers {
		var _handler Handler
		_handler = handler
		// Register middlewares.
		for i := len(r.middlewares) - 1; i >= 0; i-- {
			_handler = r.middlewares[i].Handle(_handler)
		}
		// Add to route.
		r.Router.Handle(method, route, func(_ctx *fasthttp.RequestCtx, ps router.Params) {
			ctx := NewContext(r, _ctx, &ps)
			defer ctx.Close()
			_handler.Handle(ctx)
		})
	}
}
