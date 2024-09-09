package token

import (
	"fmt"
	"time"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/requests"
	jwt "github.com/golang-jwt/jwt/v5"
)

var key = "butterfly"

func GenerateToken(id int32, role string) (*requests.UserLoginResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   id,
			"role": role,
			"exp":  time.Now().Add(time.Minute * 30).Unix(),
		})

	accessToken, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token :%v", err)
	}
	return &requests.UserLoginResponse{Token: accessToken}, nil
}

func VerifyToken(token, key string) error {
	tokens, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	if !tokens.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
