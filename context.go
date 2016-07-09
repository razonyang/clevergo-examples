// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/headwindfly/jwt"
	"github.com/headwindfly/mustache"
	"github.com/headwindfly/sessions"
	"github.com/valyala/fasthttp"
)

type Context struct {
	router *Router
	*fasthttp.RequestCtx
	RouterParams *fasthttprouter.Params
	Session      *sessions.Session
	Token        *jwt.Token // JSON WEB TOKEN
}

func NewContext(r *Router, ctx *fasthttp.RequestCtx, rps *fasthttprouter.Params) *Context {
	return &Context{
		router:       r,
		RequestCtx:   ctx,
		RouterParams: rps,
	}
}

func (ctx *Context) GetSession() {
	ctx.Session, _ = ctx.router.sessionStore.Get(&ctx.RequestCtx.Request.Header, "GOSESSION")
}

func (ctx *Context) SaveSession() error {
	return ctx.router.sessionStore.Save(&ctx.Response.Header, ctx.Session)
}

func (ctx *Context) JSON(v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
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
	ctx.Response.Header.Set("Content-Type", "application/javascript; charset=utf-8")
	jsonp := append(callback, "("...)
	jsonp = append(jsonp, json...)
	jsonp = append(jsonp, ")"...)
	ctx.Response.SetBody(jsonp)
}

func (ctx *Context) JSONPWithCode(code int, v interface{}, callback []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.JSONP(v, callback)
}

func (ctx *Context) XML(v interface{}, header string) {
	xmlBytes, err := xml.MarshalIndent(v, "", `   `)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}

	if len(header) == 0 {
		header = xml.Header
	}

	var bytes []byte
	bytes = append(bytes, header...)
	bytes = append(bytes, xmlBytes...)

	ctx.Response.Header.Set("Content-Type", "application/xml; charset=utf-8")
	ctx.Response.SetBody(bytes)
}

func (ctx *Context) XMLWithCode(code int, v interface{}, header string) {
	ctx.Response.SetStatusCode(code)
	ctx.XML(v, header)
}

func (ctx *Context) HTML(body string) {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
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

func (ctx *Context) Render(file string, args ...interface{}) {
	ctx.HTML(mustache.RenderFile(file, args...))
}

func (ctx *Context) RenderInLayout(filename string, layoutFile string, args ...interface{}) {
	ctx.HTML(mustache.RenderInLayout(filename, layoutFile, args...))
}

func (ctx *Context) ResponseForbidden(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusForbidden)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusForbidden)
	}
	ctx.Response.SetBodyString(msg)
}

func (ctx *Context) ResponseNotFound(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusNotFound)
	}
	ctx.Response.SetBodyString(msg)
}

func (ctx *Context) ResponseMethodNotAllowed(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed)
	}
	ctx.Response.SetBodyString(msg)
}

func (ctx *Context) ResponseInternalServerError(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusInternalServerError)
	}
	ctx.Response.SetBodyString(msg)
}

func (ctx *Context) ResponseUnauthorized(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusUnauthorized)
	}
	ctx.Response.SetBodyString(msg)
}

func (ctx *Context) ResponseBadRequest(args ...string) {
	ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
	var msg string
	if len(args) > 0 {
		msg = args[0]
	} else {
		msg = fasthttp.StatusMessage(fasthttp.StatusBadRequest)
	}
	ctx.Response.SetBodyString(msg)
}
