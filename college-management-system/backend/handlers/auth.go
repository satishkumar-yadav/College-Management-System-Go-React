package handlers

import (
	"cms/db"
	"cms/models"
	"cms/utils"
	"context"
	"encoding/json"
	"net/http"
)

// ⏳ Register to MongoDB (temporary storage)
func RegisterTempUser(w http.ResponseWriter, r *http.Request) {
	var user models.TempUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	/*
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		count, _ := db.TempUserCollection.CountDocuments(ctx, map[string]interface{}{"email": user.Email})
		if count > 0 {
			http.Error(w, "User already registered", http.StatusConflict)
			return
		}*/

	// Check for duplicate
	var exists bool
	row := db.MySQLDB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", user.Email)
	_ = row.Scan(&exists)

	if exists {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// _, err := db.TempUserCollection.InsertOne(ctx,user)
	_, err := db.TempUserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to save temp user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Registration submitted for approval"})
	// Registration successful, pending approval
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

	// Generate Token on Login
	token, _ := utils.GenerateJWT(user)
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
