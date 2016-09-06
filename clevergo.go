// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

const (
	// Version of CleverGo.
	Version = "1.1.1"

	// Logo of CleverGo.
	Logo = `  ____ _     _______     _______ ____   ____  ___
 / ___| |   | ____\ \   / / ____|  _ \ / ___|/ _ \
| |   | |   |  _|  \ \ / /|  _| | |_) | |  _| | | |
| |___| |___| |___  \ V / | |___|  _ <| |_| | |_| |
 \____|_____|_____|  \_/  |_____|_| \_\\____|\___/ `
)

func info() {
	fmt.Printf("\x1b[36;1m%s %s\x1b[0m\n\n\x1b[32;1mStarted at %s\x1b[0m\n", Logo, Version, time.Now())
}

// ListenAndServe is a alias of fasthttp.ListenAndServe.
func ListenAndServe(addr string, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServe(addr, handler)
}

// ListenAndServeUNIX is a alias of fasthttp.ListenAndServeUNIX.
func ListenAndServeUNIX(addr string, mode os.FileMode, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeUNIX(addr, mode, handler)
}

// ListenAndServeTLS is a alias of fasthttp.ListenAndServeTLS.
func ListenAndServeTLS(addr, certFile, keyFile string, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

// ListenAndServeTLSEmbed is a alias of fasthttp.ListenAndServeTLSEmbed.
func ListenAndServeTLSEmbed(addr string, certData, keyData []byte, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeTLSEmbed(addr, certData, keyData, handler)
}
