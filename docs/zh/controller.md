# 控制器 Controller
Controller其实是一个Middleware，它实现了中间件的Handle方法。

也许你会问为什么要将Controller设计成一个Middleware呢？其实原因很简单，
在逻辑处理之前，会有一些准备工作，以及处理之后的一些收尾工作，
如果设计成中间件，这些都可以在Handle方法里面轻松掌控。

运行这个实例可以很好的理解：[examples/controller.go](/examples/restful)

## Controller Interface
控制器定义了RESTFUL API相关的方法：GET、DELETE、POST、PUT等，参阅[controller.go](/controller.go)。

## 编写控制器
这里以上述例子的UserController为例：
```
type UserController struct {
	clevergo.Controller
}
```
UserController内嵌了**clevergo.Controller**，也就实现了Controller Interface。
如果UserController没有复写GET、POST、PUT、DELETE等方法，则默认响应为Forbidden。

路由器会根据请求的Method调用对应的方法，比如GET请求则会调用GET方法。
当然这个可以在Handle方法里改变这个策略，这个也是将控制器设计为中间件的原因之一。

## 注册控制器
你需要通过Router的RegisterController方法注册控制器，
其中第一个参数为route path,第二个参数为控制器实例。
比如router.RegisterController("/user",UserController{})

## Shortcut
* [目录](README.md)
* [路由](router.md)
* [controller.go](/controller.go)
