# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Go(go programing language).
It built on top of [**fasthttp**](https://github.com/valyala/fasthttp).

### Features
- **High performance**

    1. CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks.
    2. CleverGo's handler [**router**](https://github.com/clevergo/router) is a high performance router(fork from [fasthttprouter](https://github.com/buaazp/fasthttprouter)).
    3. Simple architecture.
    4. No reflect.

See also [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).
![Go Web Framework Benchmark](https://github.com/smallnest/go-web-framework-benchmark/raw/master/benchmark.png "Go Web Framework Benchmark")

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

### Documentation
- [**English**](https://github.com/clevergo/docs/en)
- [**中文**](https://github.com/clevergo/docs/zh)

### Relevant Projects
Most of packages can be found at https://github.com/clevergo.

- [**fasthttp**](https://github.com/valyala/fasthttp)

- [**middlewares**](https://github.com/clevergo/middlewares)
    CleverGo middlewares.

- [**router**](https://github.com/clevergo/router)

- [**websocket**](https://github.com/clevergo/websocket)

- [**sessions**](https://github.com/clevergo/sessions)

- [**captcha**](https://github.com/clevergo/captcha)

- [**csrf**](https://github.com/clevergo/csrf)

- [**jwt**](https://github.com/clevergo/jwt)

- [**mustache**](https://github.com/clevergo/mustache)

- [**utils**](https://github.com/clevergo/utils)

