// Session Middleware
//
// In order to get session.

package sessionmiddleware

import (
	"github.com/headwindfly/clevergo"
)

var (
	// Default session name.
	SessionName = "GOSESSION"
)

// Session Middleware
type SessionMiddleware struct {
	name string // Session name
}

// Returns Session Middleware instance.
func NewSessionMiddleware(name string) SessionMiddleware {
	return SessionMiddleware{
		name: name,
	}
}

// Session Middleware Handler.
//
// Implemented Middleware Interface.
func (m SessionMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		ctx.Session, _ = ctx.SessionStore().Get(ctx.RequestCtx, m.name)
		defer ctx.Session.Save(ctx.RequestCtx)

		next.Handle(ctx)
	})
}
