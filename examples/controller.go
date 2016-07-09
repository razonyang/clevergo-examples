package main

import (
	"github.com/headwindfly/clevergo"
	"log"
)

type UserController struct {
	clevergo.Controller
}

func (c UserController) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Do anything what you want.

		ctx.Text("Prepare.\n")

		// Invoke the request handler.
		next.Handle(ctx)

		ctx.Text("Finished.\n")
	})
}

func (c UserController) GET(ctx *clevergo.Context) {
	ctx.Text("GET REQUEST.\n")
}

func (c UserController) POST(ctx *clevergo.Context) {
	ctx.Text("POST REQUEST.\n")
}

func (c UserController) DELETE(ctx *clevergo.Context) {
	ctx.Text("DELETE REQUEST.\n")
}

func (c UserController) PUT(ctx *clevergo.Context) {
	ctx.Text("PUT REQUEST.\n")
}

func (c UserController) OPTIONS(ctx *clevergo.Context) {
	ctx.Text("OPTIONS REQUEST.\n")
}

func (c UserController) PATCH(ctx *clevergo.Context) {
	ctx.Text("PATCH REQUEST.\n")
}

func (c UserController) HEAD(ctx *clevergo.Context) {
	ctx.Text("HEAD REQUEST.\n")
}

func main() {
	// Create a router instance.
	router := clevergo.NewRouter()

	// Register route handler.
	router.RegisterController("/", UserController{})

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
