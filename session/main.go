package main

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/clevergo/clevergo"
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
	router := clevergo.NewRouter()
	router.Use(
		clevergo.WrapHH(sessionManager.LoadAndSave),
	)
	router.Get("/", index)
	http.ListenAndServe(":8080", router)
}
