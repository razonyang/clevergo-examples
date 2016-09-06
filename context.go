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

// Context.
type Context struct {
	router *Router
	*fasthttp.RequestCtx
	RouterParams *router.Params
	Session      *sessions.Session
}

// Returns a Context instance.
//
// Firstly, it will try to get Context instance from contextPool.
// If failed to get Context from contextPool,
// returns a new Context instance.
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

// Close Context.
//
// Context should be closed after finishing request,
// and at this moment, put the context into contextPool.
func (ctx *Context) Close() {
	ctx.Session = nil
	contextPool.Put(ctx)
}

// Returns the session store of router.
func (ctx *Context) SessionStore() sessions.Store {
	return ctx.router.sessionStore
}

// Returns logger.
//
// Returns the router's logger if the logger is non-nil.
// Otherwise, returns the default logger of ctx.
func (ctx *Context) Logger() fasthttp.Logger {
	if ctx.router.logger != nil {
		return ctx.router.logger
	}
	return ctx.RequestCtx.Logger()
}

// Set Content-Type to HTML.
func (ctx *Context) SetContentTypeToHTML() {
	ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
}

// Set Content-Type to JSON.
func (ctx *Context) SetContentTypeToJSON() {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
}

// Set Content-Type to JSONP.
func (ctx *Context) SetContentTypeToJSONP() {
	ctx.Response.Header.Set("Content-Type", "application/javascript; charset=utf-8")
}

// Set Content-Type to XML.
func (ctx *Context) SetContentTypeToXML() {
	ctx.Response.Header.Set("Content-Type", "application/xml; charset=utf-8")
}

// Response JSON data to client.
func (ctx *Context) JSON(v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		return
	}
	ctx.SetContentTypeToJSON()
	ctx.Response.SetBody(json)
}

// Response JSON data and custom status code to client.
func (ctx *Context) JSONWithCode(code int, v interface{}) {
	ctx.Response.SetStatusCode(code)
	ctx.JSON(v)
}

// Response JSONP data to client.
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

// Response JSONP data and custom status code to client.
func (ctx *Context) JSONPWithCode(code int, v interface{}, callback []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.JSONP(v, callback)
}

// Response XML data to client.
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

// Response XML data and custom status code to client.
func (ctx *Context) XMLWithCode(code int, v interface{}, headers ...string) {
	ctx.Response.SetStatusCode(code)
	ctx.XML(v, headers...)
}

// Response HTML data to client.
func (ctx *Context) HTML(body string) {
	ctx.SetContentTypeToHTML()
	ctx.Response.SetBodyString(body)
}

// Response HTML data and custom status code to client.
func (ctx *Context) HTMLWithCode(code int, body string) {
	ctx.Response.SetStatusCode(code)
	ctx.HTML(body)
}

// Response Text data to client using fmt.Fprint().
func (ctx *Context) Text(a ...interface{}) {
	fmt.Fprint(ctx, a...)
}

// Response Text data to client using fmt.Fprintf().
func (ctx *Context) Textf(format string, a ...interface{}) {
	fmt.Fprintf(ctx, format, a...)
}
