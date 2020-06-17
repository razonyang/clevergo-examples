package main

import (
	"fmt"
	"net/http"

	"clevergo.tech/clevergo"
)

func index(c *clevergo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func hello(c *clevergo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", c.Params.String("name")))
}

func main() {
	app := clevergo.New()
	app.Get("/", index)
	app.Get("/hello/:name", hello)
	app.Run(":8080")
}
