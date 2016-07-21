// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto"
	"fmt"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/jwt"
	"log"
	"math/rand"
	"strconv"
)

var (
	ttl = int64(10)
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
	// Parse the token in order to get raw token.
	token.Parse()
	ctx.HTML(fmt.Sprintf(`
	JWT token(effective within %d seconds):<br>
	<textarea rows="3" cols="60">%s</textarea><br>
	<a target="_blank" href="/verify?_jwt=%s">Verify Token</a>
	`, ttl, token.Raw.Token(), token.Raw.Token()))

}

func jwtVerify(ctx *clevergo.Context) {
	ctx.Text("Congratulation! Your token is valid.")
}

func main() {
	// Create a router instance.
	router := clevergo.NewRouter()

	// Note that.
	// Before registering middleware, we should register jwtGet handler first.
	// In order to cross over the JWT Middleware,
	// otherwise the jwtGet handler will be blocked by the JWT Middleware.
	router.GET("/", clevergo.HandlerFunc(jwtGet))

	// Add JWT Middleware
	router.AddMiddleware(clevergo.NewJWTMiddleware(j))

	// Register route handler.
	router.GET("/verify", clevergo.HandlerFunc(jwtVerify))
	router.POST("/verify", clevergo.HandlerFunc(jwtVerify))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
