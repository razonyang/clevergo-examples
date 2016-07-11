// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/headwindfly/clevergo"
	"log"
	"os"
	"path"
)

var (
	helloCleverGo = []byte("Hello CleverGo!\n")
	resourcesPath = path.Join(os.Getenv("GOPATH"), "src", "github.com", "headwindfly", "clevergo", "examples")
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

func params(ctx *clevergo.Context) {
	name := ctx.RouterParams.ByName("name")
	ctx.Textf("Hello %s.", name)
}

func multiParams(ctx *clevergo.Context) {
	param1 := ctx.RouterParams.ByName("param1")
	param2 := ctx.RouterParams.ByName("param2")
	ctx.Textf("Your params is %s and %s", param1, param2)
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
	// Navigate to http://127.0.0.1:8080/params/yourname.
	router.GET("/params/:name", clevergo.HandlerFunc(params))
	// Navigate to http://127.0.0.1:8080/multi-params/param1/param2.
	router.GET("/multi-params/:param1/:param2", clevergo.HandlerFunc(multiParams))

	// Static resource files.
	// Navigate to http://127.0.0.1:8080/examples/base.go
	router.ServeFiles("/examples/*filepath", resourcesPath)

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
