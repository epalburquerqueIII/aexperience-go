package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User Datos del usuario
type User struct {
	Email, PasswordHash, UserName string
	Role, UserID                  int
}

// https://tools.ietf.org/html/rfc7519
type TokenClaims struct {
	jwt.StandardClaims
	Role int    `json:"role"`
	Csrf string `json:"csrf"`
}

// AuthWeb Datos para Csrf y usuario en la web
type AuthWeb struct {
	CsrfSecret string
	UserName   string
}

const RefreshTokenValidTime = time.Minute * 15
const AuthTokenValidTime = time.Minute * 15
