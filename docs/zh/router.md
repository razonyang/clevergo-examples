# 路由器 Router
[router.go](/router.go)

## 创建路由器
```
router := clevergo.NewRouter()
```

## 注册Middleware
```
Router.AddMiddleware(middleware Middleware)
```

## 注册Handler
```
Router.Handle(method, path string, handler Handler)
```
* [method] GET、POST等
* [path] URL PATH， 如： "/"、"/user/:name"等
* [handler] Handler

## Session Store
```
Router.SetSessionStore(store sessions.Store)
```
运行[Session Example](/examples/session)和查看[sessions package](https://github.com/clevergo/sessions)可以快速上手。

目前仅仅支持Redis Store来存储Session，因为Session Store的设计非常简单，你可以定制自己的Store :)。

## 注意事项
* Middleware需要在Handler之前注册，否则不会起作用，这也意味着可以让某些Handler越过Middleware。

## Shortcut
* [目录](README.md)
* [上下文](context.md)