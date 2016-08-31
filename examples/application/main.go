package main

import (
	"github.com/headwindfly/clevergo"
)

func hello1(ctx *clevergo.Context) {
	ctx.HTML("hello1")
}

func hello2(ctx *clevergo.Context) {
	ctx.HTML("hello2")
}

func main() {
	// Create a application instance.
	app := clevergo.NewApplication()

	// Set configuration.
	// app.Config.ServerAddr = ":80"

	// Set default logger.
	// app.SetLogger(logger)

	// Set default session store.
	// app.SetSessionStore(store)

	// Create a router instance.
	router1 := clevergo.NewRouter()
	router2 := clevergo.NewRouter()

	router1.GET("/", clevergo.HandlerFunc(hello1))
	router2.GET("/", clevergo.HandlerFunc(hello2))

	app.AddRouter("clevergo.dev", router1)
	app.AddRouter("user.clevergo.dev", router2)

	// Change default router
	// app.SetDefaultRouter(router2)

	// Start the application.
	app.Run()
}
