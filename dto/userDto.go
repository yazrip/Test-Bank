package dto

import "github.com/golang-jwt/jwt"

type UserDto struct {
	Id       string
	Username string
	jwt.StandardClaims
}
