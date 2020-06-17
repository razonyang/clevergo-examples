package main

import (
	"fmt"

	"clevergo.tech/clevergo"
	"clevergo.tech/form"
)

var decoders = form.New()

type user struct {
	Username string `schema:"username" json:"username" xml:"username"`
	Password string `schema:"password" json:"password" xml:"password"`
}

func login(ctx *clevergo.Context) error {
	u := user{}
	if err := decoders.Decode(ctx.Request, &u); err != nil {
		return err
	}
	ctx.WriteString(fmt.Sprintf("username: %s, password: %s", u.Username, u.Password))
	return nil
}

func main() {
	app := clevergo.New()
	app.Post("/login", login)
	app.Run(":8080")
}
