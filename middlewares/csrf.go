// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"errors"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/csrf"
	"github.com/headwindfly/utils"
)

const key = "_csrf"

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
		Key:         key,
		SessionKey:  key,            // be used to save token into session.
		FormKey:     key,            // be used to get token from form's args.
		HeaderKey:   "X-CSRF-Token", // be used to get token from header.
		SafeMethods: SafeMethods,
		ErrInvalid:  ErrInvalid,
	}
}

func (m CSRFMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
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
func (m CSRFMiddleware) Token(ctx *clevergo.Context) ([]byte, error) {
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
