package config

import (
	"net/http"
	"strings"
	modelstoken "todo-api/models/models-token"
	"todo-api/server"

	"github.com/golang-jwt/jwt/v4"
)

func Token(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	token, err := jwt.ParseWithClaims(tokenString, &modelstoken.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
