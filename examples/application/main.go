package main

import (
	"github.com/headwindfly/clevergo"
)

func hello1(ctx *clevergo.Context) {
	ctx.HTML("clevergo.dev")
}

func hello2(ctx *clevergo.Context) {
	ctx.HTML("user.clevergo.dev")
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
	router1 := app.NewRouter("clevergo.dev")
	router1.GET("/", clevergo.HandlerFunc(hello1))

	router2 := app.NewRouter("user.clevergo.dev")
	router2.GET("/", clevergo.HandlerFunc(hello2))

	// Change default router
	app.SetDefaultRouter(router1)

	// Start the application.
	app.Run()
}
