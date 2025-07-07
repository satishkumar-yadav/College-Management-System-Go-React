package main

import (
	"cms/db"
	"cms/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitMongo()
	db.InitMySQL()

	r := mux.NewRouter()

	// Example test route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("College Backend Running"))
	})

	r.HandleFunc("/api/register-temp", handlers.RegisterTempUser).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	r.HandleFunc("/api/notices", handlers.GetNotices).Methods("GET")

	r.HandleFunc("/api/admin/temp-users", handlers.GetTempUsers).Methods("GET")
	r.HandleFunc("/api/admin/approve", handlers.ApproveUser).Methods("POST")
	r.HandleFunc("/api/admin/reject", handlers.RejectUser).Methods("POST")

	r.HandleFunc("/api/dashboard", JWTMiddleware(handlers.Dashboard)).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
