package clevergo

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}
