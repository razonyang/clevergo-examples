package main

import (
	"fmt"
	"net/http"

	"github.com/clevergo/clevergo"
)

func index(c *clevergo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func hello(c *clevergo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", c.Params.String("name")))
}

func main() {
	router := clevergo.NewRouter()
	router.Get("/", index)
	router.Get("/hello/:name", hello)
	http.ListenAndServe(":8080", router)
}
