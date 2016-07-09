# CleverGo
**CleverGo** is a **simple**, **high performance** and **secure** web framework for Golang. 
This project aims to become a powerful web development tool and 
make developer easily to build a high-performance, secure and stronger web application.

# Features
- **high performance**

CleverGo uses **fasthttp** instead of **net/http**, so it is more fast than net/http‘s frameworks,
and not only that, it uses [**fasthttprouter**](https://github.com/buaazp/fasthttprouter) as handler's router,
and it's architecture is very simple so make it run fast.

- **lightweight**

CleverGo's architecture is very simple, such as the design of the [**Middleware**](middleware.go).

- **easy to use**

We provides some examples below, see also [**Examples**](#examples).

- **variety of components**

| Name             | Description              | Usage                                         |
| :---             | :------------------      | :-------------------------------------------- |
| Middleware       | Middleware               | [**Middleware**](examples/middleware.go)      |
| Session          | Session                  | [**Session**](examples/session.go)            |
| RESTFUL API      | RESTFUL API              | [**Restful API**](examples/controller.go)     |
| CSRF Middleware  | CSRF attack protection   | [**CSRF Protection**](examples/csrf.go)       |
| Captcha          | Captcha                  | [**Captcha**](examples/captcha.go)            |

# Benchmark
See also [**Go Web Framework Benchmark**](https://github.com/headwindfly/go-web-framework-benchmark).

# Documentation
- [**中文文档**](docs/zh)
- [**English**](docs/en)

# Examples
- [**Basic Usages**](examples/base.go)
- [**Middleware**](examples/middleware.go)
- [**Session**](examples/session.go)
- [**Restful API**](examples/controller.go)
- [**CSRF Protection**](examples/csrf.go)
- [**Captcha**](examples/captcha.go)

# TODO LIST
- Add support for JSON WEB TOKEN(JWT).

# Relevant Projects
- [**fasthttp**](https://github.com/valyala/fasthttp)
- [**fasthttprouter**](https://github.com/buaazp/fasthttprouter)
- [**sessions**](https://github.com/headwindfly/sessions)
- [**captcha**](https://github.com/headwindfly/captcha)
- [**csrf**](https://github.com/headwindfly/csrf)
- [**mustache**](https://github.com/headwindfly/mustache)
- [**utils**](https://github.com/headwindfly/utils)
