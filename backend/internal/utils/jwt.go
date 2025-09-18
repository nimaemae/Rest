package utils

import (
	"errors"
	"time"

	"coffee-shop-platform/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, username, userType string, shopID *uint, cfg *config.Config) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"type":     userType,
		"shop_id":  shopID,
		"exp":      time.Now().Add(time.Duration(cfg.JWT.ExpireHours) * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

func ParseJWT(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}
