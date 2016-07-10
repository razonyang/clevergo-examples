# 处理器 Handler
Handler 是不可以或缺的业务处理部件。

## Handler 接口
```
type Handler interface {
	Handle(*Context)
}
```
CleverGo的**Handler**设计思路来自**net/http**包的Hander。

Handler需要实现Handle方法，其参数为*Context，参阅[上下文](context.md)

## HandlerFun
```
type HandlerFunc func(*Context)

func (hf HandlerFunc) Handle(ctx *Context) {
	hf(ctx)
}
```

## Shortcut
* [目录](README.md)
* [路由](router.md)
* [上下文](context.md)