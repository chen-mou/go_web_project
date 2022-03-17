package jwtTool

import (
	"errors"
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

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, errors.New("token 非法")
	}
	if tokenClaim != nil {
		claim, ok := tokenClaim.Claims.(*Claims)
		if ok && tokenClaim.Valid {
			return claim, nil
		}
	}
	return nil, errors.New("token失效或过期")
}
