// JSON WEB TOKEN MIDDLEWARE

package jwtmiddleware

import (
	"github.com/clevergo/jwt"
	"github.com/headwindfly/clevergo"
	"github.com/valyala/fasthttp"
)

const (
	jwtKey = "_jwt"
)

var (
	bearer     = []byte("BEARER ")
	errHandler = func(ctx *clevergo.Context) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	}
	successHandler = func(ctx *clevergo.Context, token *jwt.Token) {}
)

type JWTMiddleware struct {
	jwt            *jwt.JWT
	key            string
	errorHandler   clevergo.HandlerFunc
	successHandler func(ctx *clevergo.Context, token *jwt.Token)
}

func NewJWTMiddleware(jwt *jwt.JWT) JWTMiddleware {
	return JWTMiddleware{
		jwt:            jwt,
		key:            jwtKey,
		errorHandler:   errHandler,
		successHandler: successHandler,
	}
}

func (m JWTMiddleware) Key() string {
	return m.key
}

func (m *JWTMiddleware) SetKey(key string) {
	m.key = key
}

func (m *JWTMiddleware) SetErrorHandler(handler clevergo.HandlerFunc) {
	m.errorHandler = handler
}

func (m *JWTMiddleware) SetSuccessHandler(handler func(ctx *clevergo.Context, token *jwt.Token)) {
	m.successHandler = handler
}

func (m JWTMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Try to get JWT raw token from request form.
		rawToken := ctx.FormValue(m.key)
		if len(rawToken) == 0 {
			// Try to get JWT raw token from Header.
			if ah := ctx.Request.Header.Peek("Authorization"); len(ah) > 0 {
				// Should be a bearer token
				if len(ah) > 6 {
					rawToken = ah[7:]
				}
			}
		}

		// Check raw token is valid.
		if len(rawToken) == 0 {
			m.errorHandler(ctx)
			return
		}

		// Get JWT by raw token.
		var err error
		token, err := m.jwt.NewTokenByRaw(string(rawToken))
		if err != nil {
			m.errorHandler(ctx)
			return
		}

		// Validate Token.
		if err := token.Validate(); err != nil {
			m.errorHandler(ctx)
			return
		}

		// Invoke the success handler
		m.successHandler(ctx, token)

		// Validate successfully.
		next.Handle(ctx)
	})
}
