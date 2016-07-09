// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/sessions"
	"log"
	"math/rand"
	"time"
)

var router *clevergo.Router

func getSession(ctx *clevergo.Context) {
	// Get session.
	ctx.GetSession()
	defer ctx.SaveSession()

	if number, ok := ctx.Session.Values["randomNumber"]; ok {
		fmt.Fprint(ctx, fmt.Sprintf("The random number is: %d.\n", number))
		return
	}

	fmt.Fprint(ctx, "No random number.\n")
}

func setSession(ctx *clevergo.Context) {
	// Get session.
	ctx.GetSession()
	defer ctx.SaveSession()

	// Set random number.
	randomNumber := rand.Intn(100)
	ctx.Session.Values["randomNumber"] = randomNumber

	fmt.Fprint(ctx, fmt.Sprintf("The random number has been set as: %d.\n", randomNumber))
}

func main() {
	// Create a router instance.
	router = clevergo.NewRouter()

	// Create a redis pool.
	redisPool := &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: time.Duration(100) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
	defer redisPool.Close()

	// Create a redis session store.
	store := sessions.NewRedisStore(redisPool, sessions.Options{
		MaxAge: 3600 * 24 * 7, // 10 seconds.
		// Domain:".clevergo.dev",
		// HttpOnly:true,
		// Secure:true,
	})

	// Set session store.
	router.SetSessionStore(store)

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(getSession))
	router.GET("/random", clevergo.HandlerFunc(setSession))

	// Start server.
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
