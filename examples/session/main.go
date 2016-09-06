package main

import (
	"fmt"
	"github.com/clevergo/sessions"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/clevergo/middlewares/session"
	"html/template"
	"log"
	"math/rand"
)

var (
	tpl = template.Must(template.New("").Parse(`<html>
	<head><title>Session Example</title></head>
	<body>
		No random number has been set,
		please navigate to <a target="_blank" href="/random">here</a>.
	 	And then reload this page.
	 </body></html>`))
)

func getSession(ctx *clevergo.Context) {
	if number, ok := ctx.Session.Values["randomNumber"]; ok {
		fmt.Fprint(ctx, fmt.Sprintf("The random number is: %d.\n", number))
		return
	}

	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, nil)
}

func setSession(ctx *clevergo.Context) {
	// Set random number.
	randomNumber := rand.Intn(100)
	ctx.Session.Values["randomNumber"] = randomNumber

	fmt.Fprint(ctx, fmt.Sprintf("The random number has been set as: %d.\n", randomNumber))
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := clevergo.NewRouter()
	app.AddRouter("", router)

	// Create a redis session store.
	store := sessions.NewCookieStore([]byte("SecretKey"))

	// Set session store.
	router.SetSessionStore(store)

	// Add session middleware, this middleware is necessary.
	router.AddMiddleware(sessionmiddleware.NewSessionMiddleware("GOSESSION"))

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(getSession))
	router.GET("/random", clevergo.HandlerFunc(setSession))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
