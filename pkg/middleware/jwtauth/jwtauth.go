package jwtauth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mbrunos/go-hire/pkg/router"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func JwtAuthMiddleware(secret string) router.Middleware {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *router.Context) {
			tokenStr := c.HeaderParam("Authorization")
			if tokenStr == "" {
				c.SendError(http.StatusUnauthorized, errors.New("unauthorized"))
				return
			}

			token := strings.Split(tokenStr, "Bearer ")[1]

			claims := &Claims{}
			tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !tkn.Valid {
				c.SendError(http.StatusUnauthorized, errors.New("invalid token"))
				return
			}

			c.Set("user_id", claims.UserID)
			next(c)
		}
	}
}

func NewToken(secret, userId string, expiresIn time.Duration) (string, error) {
	exp := time.Now().Add(expiresIn)
	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
