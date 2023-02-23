package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
	"time"
)

//var jwtKey = []byte("my_key")

type Claims struct {
	UserName string
	jwt.StandardClaims
}

type JwtData struct {
	expire int
	key    string
	Header string
}

var jwtData *JwtData

func InitializeJWT(cfg *settings.JwtConfig) (err error) {
	jwtData = &JwtData{
		expire: cfg.Expire,
		key:    cfg.Key,
		Header: cfg.Header,
	}
	return
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	token := Claims{}
	claims, err := jwt.ParseWithClaims(tokenStr, &token, func(token *jwt.Token) (interface{}, error) {
		return jwtData.key, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return claims, &token, nil
}

func IsExpire(claims *Claims) bool {
	return time.Unix(claims.ExpiresAt, 0).Before(time.Now())
}

func ReleaseToken(userName string) (string, error) {
	now := time.Now()
	expireTime := now.Add(time.Duration(jwtData.expire) * time.Second)
	claims := Claims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtData.key)
}
