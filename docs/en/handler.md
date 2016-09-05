# Handler

### Handler Interface
```go
type Handler interface {
	Handle(*Context)
}
```

```go
type HandlerFunc func(*Context)
```

```go
func (hf HandlerFunc) Handle(ctx *Context) {
	hf(ctx)
}
```
You will find that the HandlerFunc is similar to `net/http` `HandlerFunc`.
Yup, the `Handler` and `HandlerFunc` are inspired by `net/http`.

### Create handler
```
type MyHandler struct {
}

func (h MyHandler) Handle(ctx *clevergo.Context) {
    ctx.SetBodyString("MyHandler")
}
```

### net/http HandlerFunc

```go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

### Shortcuts
- [Catalogue](../en)
- [Router](router.md)
- [Context](context.md)
- [Examples](/examples)