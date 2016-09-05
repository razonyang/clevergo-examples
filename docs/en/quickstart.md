# Quick Start

### Installation
```
go get github.com/headwindfly/clevergo
```

### Hello CleverGo
```
package main

import (
	"log"
	"github.com/headwindfly/clevergo"
)


func index(ctx *clevergo.Context) {
	ctx.SetBodyString("Hello CleverGo.")
}

func main() {
	// Create a router instance.
	router := clevergo.NewRouter()

	// Regitster route handler.
	router.GET("/", clevergo.HandlerFunc(index))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
```
And then navigate to http://127.0.0.1:8080, it will prints text message('Hello CleverGo.') to the browser.

As is shown above:

1. Firstly, we created a router instance by `router := clevergo.NewRouter()`,
the router is **the most important role**, it can be interpreted as `Application`.
See also [Router](router.md).

2. Secondly, we created a route's handler named `index`, and register it for route `/`(GET request).
About Handler and Context, please refer to [Handler](handler.md) and [Context](context.md).

3. Finally, we started the server.

### Basic Usage Example
[Basic Usage](/examples/basic)

### Shortcut
- [Router](router.md)
- [Context](context.md)
- [Handler](handler.md)
- [Middleware](middleware.md)
- [Examples](/examples)