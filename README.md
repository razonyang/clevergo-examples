# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Golang.
It built on top of [fasthttp](https://github.com/valyala/fasthttp).
This project aims to become a powerful web development tool and 
make developer easily to build a high-performance, secure and stronger web application.

**This project has just been launched, if you find a bug, [please tell me](https://github.com/headwindfly/clevergo/issues/new),
 I will fix it as soon as possible.**

**And all comments or suggestions regarding CleverGo web framework are welcome. :)**

# Features
- **high performance**

     CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks,
and not only that, it uses [**fasthttprouter**](https://github.com/buaazp/fasthttprouter) as handler's router,
and it's architecture is very simple and no reflect, so it runs fast.

- **lightweight**

    CleverGo's architecture is very simple, such as the design of the [**Middleware**](middleware.go) and [**Handler**](handler.go).

- **easy to use**

    We provides some examples below, see also [**Examples**](#examples).

- **components**

| Name                 | Description                                   | Usage                                         |
| :---                 | :---------------------------------------------| :-------------------------------------------- |
| **Middleware**       | Middleware                                    | [**Middleware**](examples/middleware.go)      |
| **Websocket**        | Websocket                                     | [**Websocket**](examples/websocket.go)        |
| **Session**          | Session                                       | [**Session**](examples/session.go)            |
| **RESTFUL API**      | RESTFUL API                                   | [**Restful API**](examples/controller.go)     |
| **CSRF Middleware**  | CSRF attack protection                        | [**CSRF Protection**](examples/csrf.go)       |
| **Captcha**          | Captcha                                       | [**Captcha**](examples/captcha.go)            |
| **JWT**              | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](examples/jwt.go)         |

# Benchmark
See also [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).

# Documentation
- [**中文**](https://github.com/clevergo/docs/zh)
- [**English**](https://github.com/clevergo/docs/en)

# Examples
https://github.com/clevergo/examples

- [**Basic Usages**](https://github.com/clevergo/examples/basic)
- [**Middleware**](https://github.com/clevergo/examples/middleware)
- [**WebSocket**](https://github.com/clevergo/examples/websocket)
- [**Session**](https://github.com/clevergo/examples/session)
- [**Restful API**](https://github.com/clevergo/examples/rest)
- [**CSRF Protection**](https://github.com/clevergo/examples/csrf)
- [**Captcha**](https://github.com/clevergo/examples/captcha)
- [**JSON WBE TOKEN**](https://github.com/clevergo/examples/jwt)

# Relevant Projects
- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**fasthttprouter**](https://github.com/buaazp/fasthttprouter)
- [**websocket**](https://github.com/clevergo/websocket)
- [**sessions**](https://github.com/clevergo/sessions)
- [**captcha**](https://github.com/clevergo/captcha)
- [**csrf**](https://github.com/clevergo/csrf)
- [**jwt**](https://github.com/clevergo/jwt)
- [**mustache**](https://github.com/clevergo/mustache)
- [**utils**](https://github.com/clevergo/utils)


# Official Website
**https://headwindfly.com**
This site powered by CleverGo, a **LIVE DEMO** of CleverGo.

It's source code can be found [**here**](https://github.com/headwindfly/headwindfly.com).