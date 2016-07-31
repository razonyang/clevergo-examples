// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}
