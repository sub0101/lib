package utility

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	Id    uint
	Role  string
	LibId uint
}

func CreateJwtToken(payload JwtPayload) (string, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Hour * 24 * 5).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
