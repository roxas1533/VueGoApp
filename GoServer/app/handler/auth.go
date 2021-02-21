package handler

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Secret Secret key
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

//CheckToken JWTトークンを検証します。
func CheckToken(r string) *jwt.Token {
	token, err := jwt.Parse(r, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", nil
		}
		return []byte(Secret), nil
	})
	if token == nil {
		return nil
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				fmt.Printf("%s is expired", r)
				return nil
			}
			fmt.Printf("%s is invalid", r)
			return nil
		}
		fmt.Printf("%s is invalid", r)
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("no claims")
		return nil
	}
	userID, ok := claims["name"].(string)
	if !ok {
		fmt.Printf("no claims")
		return nil
	}
	id, ok := claims["sub"].(string)
	if !ok {
		fmt.Printf("no claims")
		return nil
	}
	_ = userID
	_ = id
	return token
}
