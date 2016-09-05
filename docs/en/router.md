# Router
Router is **the most important role**, it can be interpreted as `Application`.
See also source code [router.go](/router.go).

### Create a router instance
```
router := clevergo.NewRouter()
```

### Middlewares
1. Route.SetMiddlewares(middlewares []Middleware)
2. Route.AddMiddleware(middleware Middleware)

**Important Note**:
You should register middlewares before registering handler,
otherwise middlewares doesn't work.
But because of this, we can allow some handlers cross over the middlewares.
For example:
```
router.GET("/login",loginHandler)

// If the user haven't login, request will be blocked by this middleware.
router.AddMiddleware(loginMiddleware)

router.GET("/",indexHandler)
router.GET("/other",otherHandler)
```
As is shown above, the `loginHandler` has no middleware, it won't be blocked.

### Register route handler
You can register route handler by the following ways:

1. Route.GET(path string, handler Handler)
2. Route.HEAD(path string, handler Handler)
3. Route.OPTIONS(path string, handler Handler)
4. Route.POST(path string, handler Handler)
5. Route.PUT(path string, handler Handler)
6. Route.PATCH(path string, handler Handler)
7. Route.DELETE(path string, handler Handler)
8. Route.Handle(method, path string, handler Handler)


### Register RESTFul Controller
Route.RegisterController(route string, c ControllerInterface)

See also [Controller](controller.md).

### Shortcut
- [Quick Start](quickstart.md)
- [Router](router.md)
- [Context](context.md)
- [Handler](handler.md)
- [Middleware](middleware.md)
- [Controller](controller.md)