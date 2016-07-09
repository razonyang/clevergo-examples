// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/clevergo/middlewares"
	"github.com/headwindfly/jwt"
	"log"
	"math/rand"
	"strconv"
)

var (
	ttl = int64(20)
	j   = jwt.NewJWT("CleverGo", ttl)
)

func init() {
	// Add HMACAlgorithm.
	hs256, err := jwt.NewHMACAlgorithm(crypto.SHA256, []byte("secrey key"))
	if err != nil {
		panic(err)
	}
	j.AddAlgorithm("HS256", hs256)
}

func jwtGet(ctx *clevergo.Context) {
	token, err := j.NewToken("HS256", "CleverGO", "audience"+strconv.Itoa(rand.Intn(100)))
	if err != nil {
		ctx.Textf("Failed to generate token: %s", err.Error())
		return
	}
	token.Parse()
	ctx.Textf("JWT token(effective within %d seconds):\n%s", ttl, token.Raw.Token())
}

func jwtVerify(ctx *clevergo.Context) {
	if ctx.Token != nil {
		if err := ctx.Token.Validate(); err == nil {
			ctx.Text("Congratulation! Your token is valid.")
			return
		}
	}

	ctx.Text("Sorry! Your token is invalid.")
}

func main() {
	// Create a router instance.
	router := clevergo.NewRouter()

	// Add JWT Middleware
	router.AddMiddleware(middleware.NewJWTMiddleware(j))

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(jwtGet))
	router.GET("/verify", clevergo.HandlerFunc(jwtVerify))
	router.POST("/verify", clevergo.HandlerFunc(jwtVerify))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
