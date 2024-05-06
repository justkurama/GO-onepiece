package utils

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("your_secret_key")
var refreshKey = []byte("your_refresh_secret_key")

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// GenerateToken creates an access token for a given username
func GenerateToken(login string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute) // Access token is valid for 15 minutes
	claims := &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// GenerateRefreshToken creates a refresh token for a given username
func GenerateRefreshToken(login string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Refresh token is valid for 24 hours
	claims := &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	return token, err
}
