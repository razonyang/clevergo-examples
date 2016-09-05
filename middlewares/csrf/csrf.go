// CSRF Middleware

package csrfmiddleware

import (
	"errors"
	"github.com/clevergo/csrf"
	"github.com/clevergo/utils"
	"github.com/headwindfly/clevergo"
	"github.com/valyala/fasthttp"
	"strings"
)

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

	ErrHandler = func(ctx *clevergo.Context) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(ErrInvalid)
	}
)

type CSRFMiddleware struct {
	len          int    // length of token.
	maskLen      int    // length mask.
	key          string // be used to save token into Context.UserValue.
	sessionKey   string // be used to save true token into session.
	formKey      string // be used to get token from form's args.
	headerKey    string // be used to get token from header.
	safeMethods  map[string]bool
	errorHandler clevergo.HandlerFunc
}

func NewCSRFMiddleware() CSRFMiddleware {
	return CSRFMiddleware{
		len:          32,
		maskLen:      8,
		key:          csrfKey,
		sessionKey:   csrfKey,
		formKey:      csrfKey,
		headerKey:    "X-CSRF-Token",
		safeMethods:  SafeMethods,
		errorHandler: ErrHandler,
	}
}

func (m CSRFMiddleware) Len() int {
	return m.len
}

func (m *CSRFMiddleware) SetLen(len int) {
	m.len = len
}

func (m CSRFMiddleware) MaskLen() int {
	return m.maskLen
}

func (m *CSRFMiddleware) SetMaskLen(len int) {
	m.maskLen = len
}

func (m CSRFMiddleware) Key() string {
	return m.key
}

func (m *CSRFMiddleware) SetKey(key string) {
	m.key = key
}

func (m CSRFMiddleware) SessionKey() string {
	return m.sessionKey
}

func (m *CSRFMiddleware) SetSessionKey(key string) {
	m.sessionKey = key
}

func (m CSRFMiddleware) HeaderKey() string {
	return m.headerKey
}

func (m *CSRFMiddleware) SetHeaderKey(key string) {
	m.headerKey = key
}

func (m CSRFMiddleware) FormKey() string {
	return m.formKey
}

func (m *CSRFMiddleware) SetFormKey(key string) {
	m.formKey = key
}

func (m *CSRFMiddleware) SetSafeMethods(methods []string) {
	m.safeMethods = map[string]bool{}
	for i := 0; i < len(methods); i++ {
		m.safeMethods[strings.ToUpper(methods[i])] = true
	}
}

func (m *CSRFMiddleware) SetErrorHandler(handler clevergo.HandlerFunc) {
	m.errorHandler = handler
}

func (m CSRFMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		trueToken, err := m.tureToken(ctx)

		_, safe := m.safeMethods[string(ctx.Method())]

		// Validate CSRF token.
		if !safe {
			if (err != nil) ||
				((csrf.Validate(m.maskLen, ctx.PostArgs().Peek(m.formKey), trueToken) != nil) &&
					(csrf.Validate(m.maskLen, ctx.Request.Header.Peek(m.headerKey), trueToken) != nil)) {
				m.errorHandler(ctx)
				return
			}
		}

		// Create a new encode token and save it.
		token := csrf.Generate(m.maskLen, trueToken)

		// Save the encoded token into UserValue.
		ctx.SetUserValue(m.key, token)

		// Save the true token
		ctx.Session.Values[m.sessionKey] = trueToken

		next.Handle(ctx)
	})
}

// Get true token from session.
// Returns non-nil error if the token does not exists or invalid.
func (m CSRFMiddleware) tureToken(ctx *clevergo.Context) ([]byte, error) {
	token, ok := ctx.Session.Values[m.sessionKey]
	if !ok || token == nil {
		return utils.RandomBytes(m.len), errTokenNotExists
	} else {
		if v, ok := token.([]byte); ok {
			return v, nil
		}
		return utils.RandomBytes(m.len), errTokenInvalid
	}
}
