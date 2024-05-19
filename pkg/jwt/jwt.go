package jwt

import (
	"time"

	"github.com/Giafn/goMoneySaveAndManage/configs"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.AppConfig.JWTSecret))
}
