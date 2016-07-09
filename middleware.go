// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import "fmt"

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}

type BaseMiddleware struct {
}

func (bm *BaseMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the next middleware.
		next.Handle(ctx)
	})
}
