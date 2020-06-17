package main

import (
	"fmt"
	"net/http"

	"clevergo.tech/clevergo"
	"github.com/alexedwards/scs/v2"
)

var (
	sessionManager = scs.New()
	countKey       = "count"
)

func index(c *clevergo.Context) error {
	ctx := c.Request.Context()
	count := sessionManager.GetInt(ctx, countKey)
	sessionManager.Put(ctx, countKey, count+1)
	return c.String(http.StatusOK, fmt.Sprintf("You visited this page %d times", count))
}

func main() {
	app := clevergo.New()
	app.Use(
		clevergo.WrapHH(sessionManager.LoadAndSave),
	)
	app.Get("/", index)
	app.Run(":8080")
}
