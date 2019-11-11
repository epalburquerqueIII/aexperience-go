package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User Datos del usuario
type User struct {
	Email, PasswordHash, Role, UserName string
	UserID                              int
}

// https://tools.ietf.org/html/rfc7519
type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const RefreshTokenValidTime = time.Minute * 15
const AuthTokenValidTime = time.Minute * 15
