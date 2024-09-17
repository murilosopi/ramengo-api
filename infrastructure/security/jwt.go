package security

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Id   int           `json:"id"`
	Type JWTClaimsType `json:"type"`
	jwt.RegisteredClaims
}

type JWTClaimsType string

const (
	Kitchen JWTClaimsType = "kitchen"
	User    JWTClaimsType = "user"
)

func GenerateTokenJWT(id int, claimsType JWTClaimsType) string {

	claims := &JWTClaims{
		id,
		claimsType,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		fmt.Println(err)
	}

	return t
}
