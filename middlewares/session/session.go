// Session Middleware
//
// In order to get session.

package sessionmiddleware

import (
	"github.com/headwindfly/clevergo"
)

var (
	// SessionName default session name.
	SessionName = "GOSESSION"
)

// SessionMiddleware in order to get session.
type SessionMiddleware struct {
	name string
}

// NewSessionMiddleware returns Session Middleware instance.
func NewSessionMiddleware(name string) SessionMiddleware {
	return SessionMiddleware{
		name: name,
	}
}

// Handle implemented Middleware Interface.
func (m SessionMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		ctx.Session, _ = ctx.SessionStore().Get(ctx.RequestCtx, m.name)
		defer ctx.Session.Save(ctx.RequestCtx)

		next.Handle(ctx)
	})
}
