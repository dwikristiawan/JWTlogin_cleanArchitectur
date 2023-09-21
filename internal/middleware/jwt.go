package middleware

import (
	"JWTLogin/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	SecretKey        = []byte("secret")
	RefreshSecretKey = []byte("refresh")
)

func generateToken(user model.Users, key []byte, expiration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = expiration

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateTokens(user model.Users) (model.Tokens, error) {
	accessToken, err := generateToken(user, SecretKey, time.Hour*24)
	if err != nil {
		return model.Tokens{}, err
	}
	refreshToken, err := generateToken(user, RefreshSecretKey, time.Hour*24*30)
	if err != nil {
		return model.Tokens{}, err
	}
	return model.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
