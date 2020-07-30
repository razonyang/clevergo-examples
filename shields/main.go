package main

import (
	"net/http"

	"clevergo.tech/clevergo"
	"clevergo.tech/shields"
)

func handler(c *clevergo.Context) error {
	badge := shields.New("hello", "world")
	badge.LabelColor = shields.ColorBlue
	if err := badge.ParseRequest(c.Request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, badge)
}

func main() {
	app := clevergo.New()
	app.Get("/", handler)
	app.Run(":8080")
}
