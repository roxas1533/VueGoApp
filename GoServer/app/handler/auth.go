package handler

import (
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

var Secret string = "qjG+ocH6KFhO6V1Ys1kXIY1VXTF7Ne/VztlPYasW/gSEyKkYHEha9auA/qr20+njG0qy3yRk+Nf+yMBEwzXNEQ=="

func getToken(username string, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["name"] = username
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString([]byte(Secret))
	return tokenString
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	},
	SigningMethod: jwt.SigningMethodHS512,
})
