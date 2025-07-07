package handlers

import (
	"cms/db"
	"context"
	"encoding/json"
	"net/http"
)

type Notice struct {
	Title string `json:"title"`
}

func GetNotices(w http.ResponseWriter, r *http.Request) {
	cur, err := db.TempUserCollection.Database().Collection("notices").Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		http.Error(w, "Error fetching notices", http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var results []Notice
	for cur.Next(context.TODO()) {
		var n Notice
		_ = cur.Decode(&n)
		results = append(results, n)
	}

	json.NewEncoder(w).Encode(results)
}
