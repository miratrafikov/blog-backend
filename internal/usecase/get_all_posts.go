package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/miratrafikov/blog_backend/internal/db"
	"github.com/miratrafikov/blog_backend/internal/entity"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Connection.Query("SELECT * FROM post")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	posts := entity.DeserializePosts(rows)
	json, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
