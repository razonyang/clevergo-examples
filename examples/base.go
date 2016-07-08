// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/headwindfly/clevergo"
	"log"
)

var (
	helloCleverGo = []byte("Hello CleverGo!\n")
)

type User struct {
	Name string `json:"name" xml:"name"`
	Team string `json:"team" xml:"team"`
}

func hello(ctx *clevergo.Context) {
	ctx.Write(helloCleverGo)
}

func html(ctx *clevergo.Context) {
	ctx.HTML("Hello CleverGo!\n")
}

func json(ctx *clevergo.Context) {
	ctx.JSON(User{
		Name: "HeadwindFly",
		Team: "CleverGo",
	})
}

func jsonp(ctx *clevergo.Context) {
	callback := ctx.FormValue("callback")
	ctx.JSONP(User{
		Name: "HeadwindFly",
		Team: "CleverGo",
	}, callback)
}

func xml(ctx *clevergo.Context) {
	ctx.XML(User{
		Name: "HeadwindFly",
		Team: "CleverGo",
	}, "")
}

func main() {
	// Create a router instance.
	router := clevergo.NewRouter()

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(hello))
	router.GET("/html", clevergo.HandlerFunc(html))
	router.GET("/json", clevergo.HandlerFunc(json))
	router.GET("/jsonp", clevergo.HandlerFunc(jsonp))
	router.GET("/xml", clevergo.HandlerFunc(xml))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
