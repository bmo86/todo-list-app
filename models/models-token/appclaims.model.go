package modelstoken

import "github.com/golang-jwt/jwt/v4"

type AppClaims struct {
	IdUser   uint `json:"id_user"`
	Position bool `json:"position"`
	jwt.StandardClaims
}
