// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

type Handler interface {
	Handle(*Context)
}

type HandlerFunc func(*Context)

func (hf HandlerFunc) Handle(ctx *Context) {
	hf(ctx)
}
