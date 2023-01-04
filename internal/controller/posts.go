package controller

import (
	"encoding/json"
	"net/http"

	"github.com/miratrafikov/blog-backend/internal/usecase"
)

func RoutePosts(w http.ResponseWriter, r *http.Request) {
	posts, err := usecase.GetAllPosts()
	if err != nil {
		handleError(err, w)

		return
	}

	if err := sendResponseWithData(posts, w); err != nil {
		handleError(err, w)

		return
	}
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func sendResponseWithData(data any, w http.ResponseWriter) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil
}
