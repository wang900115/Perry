package entity

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	UserID    uint
	SessionID uint
	jwt.StandardClaims
}
