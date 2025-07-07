package utils

import (
	"cms/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"email":  user.Email,
		"role":   user.Role,
		"course": user.Course,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
