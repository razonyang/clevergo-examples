# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Go(go programing language).
It built on top of [**fasthttp**](https://github.com/valyala/fasthttp).

**CleverGo** itself does not provide built-in complex functionality and configuration,
but some common features and components can be found at [https://github.com/clevergo](https://github.com/clevergo).

Such as:

- [**websocket**](https://github.com/clevergo/websocket)
- [**sessions**](https://github.com/clevergo/sessions)
- [**captcha**](https://github.com/clevergo/captcha)
- [**csrf**](https://github.com/clevergo/csrf)
- [**jwt**](https://github.com/clevergo/jwt)
- [**mustache**](https://github.com/clevergo/mustache)

### Documentation
- [**English**](https://github.com/clevergo/docs/tree/master/en)
- [**中文**](https://github.com/clevergo/docs/tree/master/zh)

### Features
- **High performance**

1. CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks.
2. CleverGo's handler [**router**](https://github.com/clevergo/router) is a high performance router(fork from [fasthttprouter](https://github.com/buaazp/fasthttprouter)).
3. Simple architecture.
4. No reflect.

Please refer to [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark) for getting more detail.

- **Lightweight**

CleverGo's architecture is very simple, such as the [**Middleware**](middleware.go) and [**Handler**](handler.go).

- **Easy to use**

We provides some examples below, see also [**Examples**](#examples).

- **Components and examples**

More examples can be found at https://github.com/clevergo/examples.

| Name                 | Description                                   | Usage                                                                              |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------------------|
| **Basic Usage**      | Basic Usage                                   | [**Basic Usage**](https://github.com/clevergo/examples/tree/master/basic)          |
| **Middleware**       | Middleware                                    | [**Middleware**](https://github.com/clevergo/examples/tree/master/middleware)      |
| **Websocket**        | Websocket                                     | [**Websocket**](https://github.com/clevergo/examples/tree/master/websocket)        |
| **Session**          | Session                                       | [**Session**](https://github.com/clevergo/examples/tree/master/session)            |
| **RESTFUL API**      | RESTFUL API                                   | [**Restful API**](https://github.com/clevergo/examples/tree/master/rest)           |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Protection**](https://github.com/clevergo/examples/tree/master/csrf)       |
| **Captcha**          | Captcha                                       | [**Captcha**](https://github.com/clevergo/examples/tree/master/captcha)            |
| **JSON WEB TOKEN**   | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](https://github.com/clevergo/examples/tree/master/jwt)         |

### Relevant Projects
Most of packages can be found at https://github.com/clevergo.

- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**middlewares**](https://github.com/clevergo/middlewares)
- [**router**](https://github.com/clevergo/router)
- [**websocket**](https://github.com/clevergo/websocket)
- [**sessions**](https://github.com/clevergo/sessions)
- [**captcha**](https://github.com/clevergo/captcha)
- [**csrf**](https://github.com/clevergo/csrf)
- [**jwt**](https://github.com/clevergo/jwt)
- [**mustache**](https://github.com/clevergo/mustache)
- [**utils**](https://github.com/clevergo/utils)

