package jwtTool

import (
	"github.com/dgrijalva/jwt-go"
	"project/main/module/user/entity"
)

type Claims struct {
	UUID  string
	Roles []entity.UserRole
	Data  *map[string]interface{}
	jwt.StandardClaims
}
