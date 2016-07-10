# 中间件 Middleware
中间件在CleverGo里，亦可称之为过滤器、预处理器。

它运行于业务逻辑之前，可以用于过滤非法请求，比如预防**CSRF**（Cross-site request forgery 跨站请求伪造）攻击、IP黑白名单等。

此外，还可以做预处理器，比如JWT中间件，用于获取JSON WEB TOKEN，以鉴别用户身份。

## 内置的中间件
* [CSRF Middleware](/examples/csrf.go) 用于预防CSRF攻击，该中间件基于Session，参阅[会话](session.md)。
* [JWT Middleware](/examples/jwt.go) JSON WEB TOKEN 中间件，用于获取认证身份信息。

## 中间件设计
### 中间件接口
```
type Middleware interface {
	Handle(next Handler) Handler
}
```
如上所示，CleverGo的中间件设计非常简单，只需要实现Handle方法即可。

### 案例
[Middleware Example](/examples/middleware.go)

## Shortcut
* [目录](README.md)