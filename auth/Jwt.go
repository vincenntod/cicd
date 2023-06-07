package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("vincengantengbangetnomor1sedunia")

type JWTClaim struct {
	Id       int
	Username string
	RoleId   int
	jwt.RegisteredClaims
}
