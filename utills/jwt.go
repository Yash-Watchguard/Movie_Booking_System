package utills

import (
	"errors"
	"time"

	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("yashgoyal123")

func GenerateJwt(userId string, role role.Role) (string, error) {

	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["role"] = string(role)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecret)
}

func VarifyJwt(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sifning method")
		}
		return JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
