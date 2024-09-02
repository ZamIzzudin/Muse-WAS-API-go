package utilities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	AccessTokenSecret  = []byte("your_access_token_secret")
	RefreshTokenSecret = []byte("your_refresh_token_secret")
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateAccessToken generates a new JWT access token
func GenerateAccessToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessTokenSecret)
}

// GenerateRefreshToken generates a new JWT refresh token
func GenerateRefreshToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshTokenSecret)
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}