package handlers

import (
	"cms/db"
	"cms/models"
	"cms/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTempUsers(w http.ResponseWriter, r *http.Request) {
	cur, err := db.TempUserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var users []models.TempUser
	for cur.Next(context.TODO()) {
		var u models.TempUser
		_ = cur.Decode(&u)
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

func ApproveUser(w http.ResponseWriter, r *http.Request) {
	var user models.TempUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Insert to MySQL
	_, err := db.MySQLDB.Exec(
		"INSERT INTO users (name, email, password, role, course) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Role, user.Course,
	)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	// Remove from MongoDB
	_, _ = db.TempUserCollection.DeleteOne(context.TODO(), bson.M{"email": user.Email})

	// to := user.Email
	// name := user.Name
	//Email Notification on Approval after successful MySQL insert.
	// SendApprovalEmail(to, name)
	_ = utils.SendApprovalEmail(user.Email, user.Name)

	json.NewEncoder(w).Encode(map[string]string{"message": "User approved successfully"})
}

func RejectUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
	}
	_ = json.NewDecoder(r.Body).Decode(&payload)

	_, err := db.TempUserCollection.DeleteOne(context.TODO(), bson.M{"email": payload.Email})
	if err != nil {
		http.Error(w, "Rejection failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User rejected"})
}

func SendApprovalEmail(to string, name string) error {
	auth := smtp.PlainAuth("", "satishkumar.yadav.7549@gmail.com", "your-app-password", "smtp.gmail.com")
	msg := []byte(fmt.Sprintf("Subject: Registration Approved!\r\n\r\nHi %s,\n\nYour account has been approved. You can now log in.\n\nRegards,\nXYZ College", name))

	return smtp.SendMail("smtp.gmail.com:587", auth, "your-email@gmail.com", []string{to}, msg)
}
