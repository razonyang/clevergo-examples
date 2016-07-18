# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Golang. 
This project aims to become a powerful web development tool and 
make developer easily to build a high-performance, secure and stronger web application.

**This project has just been launched, if you find a bug , please tell me,
 I will fix it as soon as possible.**

**And all comments or suggestions regarding CleverGo web framework are welcome. :)**

# Official Website
**https://headwindfly.com**
This site powered by CleverGo, a **LIVE DEMO** of CleverGo.

# Features
- **high performance**

CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks,
and not only that, it uses [**fasthttprouter**](https://github.com/buaazp/fasthttprouter) as handler's router,
and it's architecture is very simple, so it run fast.

- **lightweight**

CleverGo's architecture is very simple, such as the design of the [**Middleware**](middleware.go).

- **easy to use**

We provides some examples below, see also [**Examples**](#examples).

- **components**

| Name             | Description                                   | Usage                                         |
| :---             | :---------------------------------------------| :-------------------------------------------- |
| Middleware       | Middleware                                    | [**Middleware**](examples/middleware.go)      |
| Session          | Session                                       | [**Session**](examples/session.go)            |
| RESTFUL API      | RESTFUL API                                   | [**Restful API**](examples/controller.go)     |
| CSRF Middleware  | CSRF attack protection                        | [**CSRF Protection**](examples/csrf.go)       |
| Captcha          | Captcha                                       | [**Captcha**](examples/captcha.go)            |
| JWT              | JSON WEB TOKEN                                | [**JSON WBE TOKEN**](examples/jwt.go)         |

# Benchmark
See also [**Go Web Framework Benchmark**](https://github.com/smallnest/go-web-framework-benchmark).

# Documentation
- [**中文**](docs/zh)
- [**English**](docs/en)

# Examples
- [**Basic Usages**](examples/base.go)
- [**Middleware**](examples/middleware.go)
- [**Session**](examples/session.go)
- [**Restful API**](examples/controller.go)
- [**CSRF Protection**](examples/csrf.go)
- [**Captcha**](examples/captcha.go)
- [**JSON WBE TOKEN**](examples/jwt.go)

# TODO LIST
- **Documentation**
- **WebSocket**

# Relevant Projects
- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**fasthttprouter**](https://github.com/buaazp/fasthttprouter)
- [**sessions**](https://github.com/headwindfly/sessions)
- [**captcha**](https://github.com/headwindfly/captcha)
- [**csrf**](https://github.com/headwindfly/csrf)
- [**jwt**](https://github.com/headwindfly/jwt)
- [**mustache**](https://github.com/headwindfly/mustache)
- [**utils**](https://github.com/headwindfly/utils)
