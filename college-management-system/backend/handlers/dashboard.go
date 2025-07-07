package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	/*
		auth := r.Header.Get("Authorization")
		tokenStr := auth[len("Bearer "):]
		claims := jwt.MapClaims{}   */
	tokenStr := r.Header.Get("Authorization")[len("Bearer "):]
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
		// return []byte("supersecretkey"), nil
	})

	email := claims["email"].(string)
	role := claims["role"].(string)
	course := claims["course"].(string)

	//json.NewEncoder(w).Encode(claims)
	json.NewEncoder(w).Encode(map[string]string{
		"email":  email,
		"role":   role,
		"course": course,
	})
}
