package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"ramengo/infrastructure/security"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	UserJWTMiddleware    echo.MiddlewareFunc = jwtMiddleware(security.User)
	KitchenJWTMiddleware echo.MiddlewareFunc = jwtMiddleware(security.Kitchen)
)

// create a typed claims middleware
func jwtMiddleware(typeMiddleware security.JWTClaimsType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get token from header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			// extract token prefix
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // validate prefix
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			// token parse token string to custom struct claims
			token, err := jwt.ParseWithClaims(tokenString, &security.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				// validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			// validate claims type token
			claims, ok := token.Claims.(*security.JWTClaims)
			if !ok || !token.Valid || claims.Type != typeMiddleware {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			// verify expiration
			if claims.ExpiresAt.Time.Before(time.Now()) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token expired")
			}

			// add claims into its type key context
			c.Set(string(claims.Type), claims)

			return next(c) // call next handler
		}
	}
}
