// Session Middleware
//
// In order to get session.

package sessionmiddleware

import (
	"github.com/headwindfly/clevergo"
)

var (
	// SessionName is the default session name.
	SessionName = "GOSESSION"
)

// SessionMiddleware
type SessionMiddleware struct {
	name string // Session name
}

// Returns SessionMiddleware instance.
func NewSessionMiddleware(name string) SessionMiddleware {
	return SessionMiddleware{
		name: name,
	}
}

// SessionMiddleware Handler.
//
// Implemented Middleware Interface.
func (m SessionMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		ctx.Session, _ = ctx.SessionStore().Get(ctx.RequestCtx, m.name)
		defer ctx.Session.Save(ctx.RequestCtx)

		next.Handle(ctx)
	})
}
