package app

import (
	"time"
	"my-service/global"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId uint32 `json:"userId"`
	jwt.StandardClaims
}

func GetJWTSecret() any {
	return []byte(global.JWTSettings.AppSecret)
}

func GenerateToken(uid uint32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSettings.Expire)

	claims := Claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSettings.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
