package handlers

import (
	"cms/db"
	"cms/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RegisterTempUser(w http.ResponseWriter, r *http.Request) {
	var user models.TempUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Check for duplicate
	var exists bool
	row := db.MySQLDB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", user.Email)
	_ = row.Scan(&exists)

	if exists {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	_, err := db.TempUserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to save temp user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Registration submitted for approval"})
}

// Generate Token on Login
var jwtKey = []byte("your-secret-key")

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

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	_ = json.NewDecoder(r.Body).Decode(&req)

	var user models.User
	err := db.MySQLDB.QueryRow("SELECT id, name, email, password, role, course FROM users WHERE email=? AND password=?", req.Email, req.Password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Course)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	// json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(map[string]string{
		"token":  token,
		"name":   user.Name,
		"role":   user.Role,
		"course": user.Course,
	})

}
