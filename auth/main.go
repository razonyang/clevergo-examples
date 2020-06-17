package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"clevergo.tech/auth"
	"clevergo.tech/auth/authenticators"
	"clevergo.tech/authmidware"
	"clevergo.tech/clevergo"
)

var authenticator auth.Authenticator

func init() {
	store := &store{
		users: []user{
			{"foo", "footoken"},
			{"bar", "bartoken"},
		},
	}
	authenticator = authenticators.NewComposite(
		authenticators.NewQueryToken(store),
		authenticators.NewBearerToken(store),
	)
}

func main() {
	app := clevergo.New()
	app.Use(authmidware.New(authenticator))
	app.Get("/auth", func(c *clevergo.Context) error {
		identity := auth.GetIdentity(c.Context())
		if identity == nil {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		return c.String(http.StatusOK, fmt.Sprintf("hello %s", identity.GetID()))
	})
	app.Run(":8080")
}

type user struct {
	id    string
	token string
}

func (u user) GetID() string {
	return u.id
}

type store struct {
	users []user
}

var errNoUser = errors.New("user does not exist")

func (s *store) GetIdentity(ctx context.Context, id string) (auth.Identity, error) {
	for _, u := range s.users {
		if u.id == id {
			return u, nil
		}
	}

	return nil, errNoUser
}

func (s *store) GetIdentityByToken(ctx context.Context, token, tokenType string) (auth.Identity, error) {
	for _, u := range s.users {
		if u.token == token {
			return u, nil
		}
	}

	return nil, errNoUser
}
