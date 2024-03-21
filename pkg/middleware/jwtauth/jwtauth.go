package jwtauth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mbrunos/go-hire/pkg/router"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.Claims
}

func JwtAuth(secret string) router.Middleware {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *router.Context) {
			token := c.HeaderParam("Authorization")
			if token == "" {
				c.SendError(401, errors.New("unauthorized"))
				return
			}

			claims := &Claims{}
			tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				c.SendError(401, errors.New("unauthorized"))
				return
			}
			if !tkn.Valid {
				c.SendError(401, errors.New("invalid token"))
				return
			}

			c.Set("claims", claims)
			println("claims", claims)
			next(c)
		}
	}
}
