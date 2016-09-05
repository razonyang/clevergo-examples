# Context
```
type Context struct {
	router *Router
	*fasthttp.RequestCtx
	RouterParams *router.Params
	Session      *sessions.Session
	Token        *jwt.Token // JSON WEB TOKEN
}
```
The Context is extension of [**fasthttp.RequestCtx**](https://github.com/valyala/fasthttp).

See also source code [context.go](/context.go).

For easy to use, it provides some useful methods:

1. Context.JSON(v interface{})
2. Context.JSONWithCode(code int, v interface{})
3. Context.JSONP(v interface{}, callback []byte)
4. Context.JSONPWithCode(code int, v interface{}, callback []byte)
5. Context.XML(v interface{}, headers ...string)
6. Context.XMLWithCode(code int, v interface{}, headers ...string)
7. Context.HTML(body string)
8. Context.HTMLWithCode(code int, body string)
9. Context.Text(a ...interface{})
10. Context.Textf(format string, a ...interface{})
11. Context.Render(file string, args ...interface{})
12. Context.RenderInLayout(filename string, layoutFile string, args ...interface{})
13. Context.ResponseForbidden(args ...string)
14. Context.ResponseNotFound(args ...string)
15. Context.ResponseMethodNotAllowed(args ...string)
16. Context.ResponseInternalServerError(args ...string)
17. Context.ResponseUnauthorized(args ...string)
18. Context.ResponseBadRequest(args ...string)

### net/http and fasthttp
| net/http                       | fasthttp                                                                      |
| :------------------------------| :-----------------------------------------------------------------------------|
|r.Body                          |ctx.PostBody()                                                                 |
|r.URL.Path                      |ctx.Path()                                                                     |
|r.URL                           |ctx.URI()                                                                      |
|r.Method                        |ctx.Method()                                                                   |
|r.Header                        |ctx.Request.Header                                                             |
|r.Header.Get()                  |ctx.Request.Header.Peek()                                                      |
|r.Host                          |ctx.Host()                                                                     |
|r.Form                          |ctx.QueryArgs() + ctx.PostArgs()                                               |
|r.PostForm                      |ctx.PostArgs()                                                                 |
|r.FormValue()                   |ctx.FormValue()                                                                |
|r.FormFile()                    |ctx.FormFile()                                                                 |
|r.MultipartForm                 |ctx.MultipartForm()                                                            |
|r.RemoteAddr                    |ctx.RemoteAddr()                                                               |
|r.RequestURI                    |ctx.RequestURI()                                                               |
|r.TLS                           |ctx.IsTLS()                                                                    |
|r.Cookie()                      |ctx.Request.Header.Cookie()                                                    |
|r.Referer()                     |ctx.Referer()                                                                  |
|r.UserAgent()                   |ctx.UserAgent()                                                                |
|w.Header()                      |ctx.Response.Header                                                            |
|w.Header().Set()                |ctx.Response.Header.Set()                                                      |
|w.Header().Set("Content-Type")  |ctx.SetContentType()                                                           |
|w.Header().Set("Set-Cookie")    |ctx.Response.Header.SetCookie()                                                |
|w.Write()                       |ctx.Write(), ctx.SetBody(), ctx.SetBodyStream(), ctx.SetBodyStreamWriter()     |
|w.WriteHeader()                 |ctx.SetStatusCode()                                                            |
|w.(http.Hijacker).Hijack()      |ctx.Hijack()                                                                   |
|http.Error()                    |ctx.Error()                                                                    |
|http.FileServer()               |fasthttp.FSHandler(), fasthttp.FS                                              |
|http.ServeFile()                |fasthttp.ServeFile()                                                           |
|http.Redirect()                 |ctx.Redirect()                                                                 |
|http.NotFound()                 |ctx.NotFound()                                                                 |

### Shortcuts
- [Catalogue](../en)
- [Examples](/examples)