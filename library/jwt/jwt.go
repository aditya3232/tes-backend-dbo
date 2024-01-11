package jwt

import (
	"errors"
	"time"

	"github.com/aditya3232/tes-backend-dbo/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte(config.CONFIG.JWT_KEY)

func GenerateToken(userID int, expired int) (string, error) {
	claims := jwt.MapClaims{
		"user": gin.H{
			"id": userID,
		},
	}

	if expired != 0 {
		claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(expired)).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 365 * 100).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil

}

func GetUserIDFromToken(encodedToken string) (int, error) {
	token, err := ValidateToken(encodedToken)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user"].(map[string]interface{})["id"].(float64)
		if !ok {
			return 0, errors.New("invalid token claims")
		}

		return int(userID), nil
	} else {
		return 0, errors.New("invalid token claims")
	}
}
