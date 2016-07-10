# 上下文 Context
Context上下文包含了Request和Response，以及一些常用的方法

## 成员
* [fasthttp.RequestCtx]的所有成员
* [RouterParams] Router参数，参阅[Router](router.md)，ByName(name string)用于获取PATH的参数，比如"/user/:name"，对应RouterParams.ByName("name")
* [Session] Session，参阅[Session](session.md)
* [Token] JSON WEB TOKEN

## 方法
* [fasthttp.RequestCtx]的所有方法
* [JSON]
* [JSONWithCode]
* [JSONP]
* [JSONPWithCode]
* [XML]
* [XMLWithCode]
* [HTML]
* [HTMLWithCode]
* [Text]
* [Textf]
* [Render]
* [RenderInLayout]
* [ResponseForbidden]
* [ResponseNotFound]
* [ResponseMethodNotAllowed]
* [ResponseInternalServerError]
* [ResponseUnauthorized]
* [ResponseBadRequest]

## net/http和fasthttp对照表
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

## Shortcut
* [目录](README.md)
* [路由](router.md)
* [处理器](handler.md)