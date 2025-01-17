package jwtUtils

import (
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("true courage is about knowing not when to take a life, but when to spare one")

type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
