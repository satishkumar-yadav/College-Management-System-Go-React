package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")[len("Bearer "):]
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	email := claims["email"].(string)
	role := claims["role"].(string)
	course := claims["course"].(string)

	json.NewEncoder(w).Encode(map[string]string{
		"email":  email,
		"role":   role,
		"course": course,
	})
}
