package jwtTool

import (
	"github.com/dgrijalva/jwt-go"
	"project/main/tool/encryption"
	"time"
)

var key []byte = []byte(encryption.MD5("web_server_token"))

func GetToken(UUID string, value *map[string]interface{}) string {
	claim := Claims{
		UUID: UUID,
		Data: value,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(key)
	if err != nil {
		panic(any(err))
	}
	return token
}
