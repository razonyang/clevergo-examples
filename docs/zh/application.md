# 应用管理器 Application
**Application**用于管理多个**Router**，而一个**Router**可理解为一个小应用，**Application**则可理解为应用管理器。

## 公开的接口
- 创建Application实例
```
app := clevergo.NewApplication()
```

- 设置默认的Logger
```
app.SetLogger(logger)
```

- 设置默认的Session Store
```
app.SetSessionStore(store)
```

- 将路由器分配给Application管理
```
app.AddRouter("domain",router)
```
第一个参数是域名，Application将会根据域名，然后将请求分配给相应的路由器。
如果第一个参数是空字符串，则会将该路由器自动设置为默认路由器

- 设置默认的路由器
```
app.SetDefaultRouter(router)
```

- 启动Application
```
app.Run()
```


## 举个例子
现有三个应用：1.clevergo.dev、2.user.clevergo.dev、3.api.clevergo.dev，这三个需要占用同一个端口（80）。
```
// 首先创建一个 Application
app := clevergo.NewApplication()

// clevergo.dev 路由器
router1 := clevergo.NewRouter()

// user.clevergo.dev 路由器
router2 := clevergo.NewRouter()

// api.clevergo.dev 路由器
router3 := clevergo.NewRouter()

// 将三个路由器分配给 Application 管理
app.AddRouter("",router1) // 这里的第一个参数为空，则将其设置为默认路由器
app.AddRouter("user.clevergo.dev",router2)
app.AddRouter("api.clevergo.dev",router3)

// 启动 Application
app.Run()
```


## Shortcut
* [目录](README.md)
* [应用](application.md)
* [上下文](context.md)
* [路由器](router.md)
* [中间件](middleware.md)