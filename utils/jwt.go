package utils

import (
	"blog-backend/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("my_secret_key")

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"userName"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	mapClaims := jwt.MapClaims{
		"userID":   userID,
		"userName": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(config.GetSecretKey()))
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return []byte(config.GetSecretKey()), nil
	//})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
