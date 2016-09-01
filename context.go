// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/clevergo/router"
	"github.com/clevergo/sessions"
	"github.com/valyala/fasthttp"
	"sync"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

type Context struct {
	router *Router
	*fasthttp.RequestCtx
	RouterParams *router.Params
	Session      *sessions.Session
}

func NewContext(r *Router, ctx *fasthttp.RequestCtx, rps *router.Params) *Context {
	if context, ok := contextPool.Get().(*Context); ok {
		context.router = r
		context.RequestCtx = ctx
		context.RouterParams = rps
		return context
	}

	return &Context{
		router:       r,
		RequestCtx:   ctx,
		RouterParams: rps,
	}
}

func (ctx *Context) Close() {
	ctx.RouterParams = nil
	ctx.RequestCtx = nil
	ctx.Session = nil
	contextPool.Put(ctx)
}

func (ctx *Context) GetSession() {
	ctx.Session, _ = ctx.router.sessionStore.Get(ctx.RequestCtx, "GOSESSION")
}

func (ctx *Context) SaveSession() error {
	return ctx.Session.Save(ctx.RequestCtx)
}

func (ctx *Context) Logger() fasthttp.Logger {
	if ctx.router.logger != nil {
		return ctx.router.logger
	}
	return ctx.RequestCtx.Logger()
}

func (ctx *Context) SetContentTypeToHTML() {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
}

func (ctx *Context) SetContentTypeToJSON() {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
}

func (ctx *Context) SetContentTypeToJSONP() {
	ctx.Response.Header.Set("Content-Type", "application/javascript; charset=utf-8")
}

func (ctx *Context) SetContentTypeToXML() {
	ctx.Response.Header.Set("Content-Type", "application/xml; charset=utf-8")
}

func (ctx *Context) JSON(v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.SetContentTypeToJSON()
	ctx.Response.SetBody(json)
}

func (ctx *Context) JSONWithCode(code int, v interface{}) {
	ctx.Response.SetStatusCode(code)
	ctx.JSON(v)
}

func (ctx *Context) JSONP(v interface{}, callback []byte) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.SetContentTypeToJSONP()
	jsonp := append(callback, "("...)
	jsonp = append(jsonp, json...)
	jsonp = append(jsonp, ")"...)
	ctx.Response.SetBody(jsonp)
}

func (ctx *Context) JSONPWithCode(code int, v interface{}, callback []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.JSONP(v, callback)
}

func (ctx *Context) XML(v interface{}, headers ...string) {
	xmlBytes, err := xml.MarshalIndent(v, "", `   `)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}

	header := xml.Header
	if len(headers) > 0 {
		header = headers[0]
	}

	var bytes []byte
	bytes = append(bytes, header...)
	bytes = append(bytes, xmlBytes...)

	ctx.SetContentTypeToXML()
	ctx.Response.SetBody(bytes)
}

func (ctx *Context) XMLWithCode(code int, v interface{}, headers ...string) {
	ctx.Response.SetStatusCode(code)
	ctx.XML(v, headers...)
}

func (ctx *Context) HTML(body string) {
	ctx.SetContentTypeToHTML()
	ctx.Response.SetBodyString(body)
}

func (ctx *Context) HTMLWithCode(code int, body string) {
	ctx.Response.SetStatusCode(code)
	ctx.HTML(body)
}

func (ctx *Context) Text(a ...interface{}) {
	fmt.Fprint(ctx, a...)
}

func (ctx *Context) Textf(format string, a ...interface{}) {
	fmt.Fprintf(ctx, format, a...)
}
