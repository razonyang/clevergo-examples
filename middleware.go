// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import "fmt"

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}

type FirstMiddleware struct {
}

func (fm *FirstMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		fmt.Println("1\n")
		ctx.Response.Header.Add("FirstMiddleware", "FirstMiddleware")
		next.Handle(ctx)
	})
}
