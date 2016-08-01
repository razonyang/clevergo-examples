# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Golang.
It built on top of [fasthttp](https://github.com/valyala/fasthttp).
This project aims to become a powerful web development tool and 
make developer easily to build a high-performance, secure and stronger web application.

# Features
- **high performance**
    CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks,
and not only that, it uses [**router**](https://github.com/buaazp/fasthttprouter) as handler's router,
and it's architecture is very simple and no reflect, so it runs fast.

- **lightweight**
    CleverGo's architecture is very simple, such as the design of the [**Middleware**](middleware.go) and [**Handler**](handler.go).

- **easy to use**
    We provides some examples below, see also [**Examples**](#examples).

- **components**

| Name                 | Description                                   | Usage                                                                  |
| :---                 | :---------------------------------------------| :----------------------------------------------------------------------|
| **Middleware**       | Middleware                                    | [**Middleware**](https://github.com/clevergo/examples/tree/master/middleware)      |
| **Websocket**        | Websocket                                     | [**Websocket**](https://github.com/clevergo/examples/tree/master/websocket)        |
| **Session**          | Session                                       | [**Session**](https://github.com/clevergo/examples/tree/master/session)            |
| **RESTFUL API**      | RESTFUL API                                   | [**Restful API**](https://github.com/clevergo/examples/tree/master/rest)           |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Protection**](https://github.com/clevergo/examples/tree/master/csrf)       |
| **Captcha**          | Captcha                                       | [**Captcha**](https://github.com/clevergo/examples/tree/master/captcha)            |
| **JSON WEB TOKEN**   | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](https://github.com/clevergo/examples/tree/master/jwt)         |

# Benchmark
See also [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).
![Go Web Framework Benchmark](https://github.com/smallnest/go-web-framework-benchmark/raw/master/benchmark.png "Go Web Framework Benchmark")

# Documentation
- [**English**](https://github.com/clevergo/docs/en)
- [**中文**](https://github.com/clevergo/docs/zh)

# Examples
https://github.com/clevergo/examples

- [**Basic Usages**](https://github.com/clevergo/examples/tree/master/basic)
- [**Middleware**](https://github.com/clevergo/examples/tree/master/middleware)
- [**WebSocket**](https://github.com/clevergo/examples/tree/master/websocket)
- [**Session**](https://github.com/clevergo/examples/tree/master/session)
- [**Restful API**](https://github.com/clevergo/examples/tree/master/rest)
- [**CSRF Protection**](https://github.com/clevergo/examples/tree/master/csrf)
- [**Captcha**](https://github.com/clevergo/examples/tree/master/captcha)
- [**JSON WBE TOKEN**](https://github.com/clevergo/examples/tree/master/jwt)

# Relevant Projects
Most of packages can be found at https://github.com/clevergo

- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**router**](https://github.com/clevergo/router)
- [**websocket**](https://github.com/clevergo/websocket)
- [**sessions**](https://github.com/clevergo/sessions)
- [**captcha**](https://github.com/clevergo/captcha)
- [**csrf**](https://github.com/clevergo/csrf)
- [**jwt**](https://github.com/clevergo/jwt)
- [**mustache**](https://github.com/clevergo/mustache)
- [**utils**](https://github.com/clevergo/utils)
