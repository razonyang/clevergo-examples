# Middleware
[Middleware](https://en.wikipedia.org/wiki/Middleware) in [CLeverGo](https://github.com/headwindfly/clevergo) can be used as **blocker**, **filter** or the **preprocessor**.

### Middleware Interface
```
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}
```
About the **Handler**, please see also [Handler](handler.md).

### Create a middleware
As is shown above, a middleware need to Implement the method named `Handle(next clevergo.Handler) clevergo.Handler`.

See also [**Middleware Example**](/examples/middleware).

### Shortcuts
- [Catalogue](../en)
- [Handler](handler.md)
- [Examples](/examples)