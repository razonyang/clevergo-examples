// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"bytes"
	"errors"
	"github.com/headwindfly/csrf"
	"github.com/headwindfly/jwt"
	"github.com/headwindfly/utils"
)

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}

type BaseMiddleware struct {
}

func (bm *BaseMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the next middleware.
		next.Handle(ctx)
	})
}

// JSON WEB TOKEN Middleware
const (
	jwtUrlKey  = "_jwt"
	jwtformKey = "_jwt"
)

var (
	bearer = []byte("BEARER ")
)

type JWTMiddleware struct {
	jwt     *jwt.JWT
	urlKey  string
	formKey string
}

func NewJWTMiddleware(jwt *jwt.JWT) JWTMiddleware {
	return JWTMiddleware{
		jwt:     jwt,
		urlKey:  jwtUrlKey,
		formKey: jwtformKey,
	}
}

func (jm JWTMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Try to get JWT raw token from URL query string.
		rawToken := ctx.FormValue(jm.urlKey)
		if len(rawToken) == 0 {
			// Try to get JWT raw token from POST FORM.
			rawToken = ctx.FormValue(jm.formKey)
			if len(rawToken) == 0 {
				// Try to get JWT raw token from Header.
				if ah := ctx.Request.Header.Peek("Authorization"); len(ah) > 0 {
					// Should be a bearer token
					if len(ah) > 6 && bytes.Equal(ah[:7], bearer) {
						rawToken = ah[7:]
					}
				}
			}
		}

		// Check raw token is valid.
		if len(rawToken) == 0 {
			ctx.ResponseUnauthorized()
			return
		}

		// Get JWT by raw token.
		var err error
		ctx.Token, err = jm.jwt.NewTokenByRaw(string(rawToken))
		if err != nil {
			ctx.ResponseUnauthorized()
			return
		}

		// Validate Token.
		if err := ctx.Token.Validate(); err != nil {
			ctx.ResponseUnauthorized()
			return
		}

		// Validate successfully.
		next.Handle(ctx)
	})
}

// CSRF Middleware
const csrfKey = "_csrf"

var (
	ErrInvalid        = "Unable to verify your data submission."
	errTokenNotExists = errors.New("The token does not exists.")
	errTokenInvalid   = errors.New("The token is invalid.")
	SafeMethods       = map[string]bool{
		"GET":     true,
		"HEAD":    true,
		"OPTIONS": true,
		"TRACE":   true,
	}
)

type CSRFMiddleware struct {
	Len         int
	Key         string
	SessionKey  string
	FormKey     string
	HeaderKey   string
	MaskLen     int
	SafeMethods map[string]bool
	ErrInvalid  string
}

func NewCSRFMiddleware() CSRFMiddleware {
	return CSRFMiddleware{
		Len:         32, // length of token.
		MaskLen:     8,  // length mask.
		Key:         csrfKey,
		SessionKey:  csrfKey,        // be used to save token into session.
		FormKey:     csrfKey,        // be used to get token from form's args.
		HeaderKey:   "X-CSRF-Token", // be used to get token from header.
		SafeMethods: SafeMethods,
		ErrInvalid:  ErrInvalid,
	}
}

func (m CSRFMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		if ctx.Session == nil {
			ctx.GetSession()
			defer ctx.SaveSession()
		}

		trueToken, err := m.Token(ctx)

		_, safe := m.SafeMethods[string(ctx.Method())]

		// Validate CSRF token.
		if !safe {
			if (err != nil) ||
				((csrf.Validate(m.MaskLen, ctx.FormValue(m.FormKey), trueToken) != nil) &&
					(csrf.Validate(m.MaskLen, ctx.Request.Header.Peek(m.HeaderKey), trueToken) != nil)) {
				ctx.ResponseBadRequest(m.ErrInvalid)
				return
			}
		}

		// Create a new encode token and save it.
		token := csrf.Generate(m.MaskLen, trueToken)
		ctx.SetUserValue(m.Key, token)
		ctx.Session.Set(m.SessionKey, trueToken)

		next.Handle(ctx)
	})
}

// Get token from session.
// Returns non-nil error if the token does not exists or invalid.
func (m CSRFMiddleware) Token(ctx *Context) ([]byte, error) {
	token, err := ctx.Session.Get(m.SessionKey)
	if (err != nil) || (token == nil) {
		return utils.RandomBytes(m.Len), errTokenNotExists
	} else {
		if v, ok := token.([]byte); ok {
			return v, nil
		}
		return utils.RandomBytes(m.Len), errTokenInvalid
	}
}
