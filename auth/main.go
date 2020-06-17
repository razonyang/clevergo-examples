package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"clevergo.tech/auth"
	"clevergo.tech/auth/authenticators"
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
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		identity := auth.GetIdentity(r.Context())
		if identity == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		io.WriteString(w, fmt.Sprintf("hello %s", identity.GetID()))
	})
	http.ListenAndServe(":8080", auth.NewMiddleware(authenticator)(http.DefaultServeMux))
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
