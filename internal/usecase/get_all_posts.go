package usecase

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/miratrafikov/blog_backend/internal/db"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	posts, err := db.Connection.Query("SELECT * FROM post;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting the posts"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postsToPseudoJson(posts)))
}

func postsToPseudoJson(posts *sql.Rows) string {
	postElements := ""
	i := 0
	for posts.Next() {
		var id int
		var title string
		var content string
		var createdAt time.Time
		posts.Scan(&id, &title, &content, &createdAt)
		if i > 0 {
			postElements += ",\n"
		}

		postElements += fmt.Sprintf("  {\n    id: %d,\n    title: %s,\n    content: %s,\n    createdAt: %d\n  }", id, title, content, createdAt.Unix())
		i++
	}

	return fmt.Sprintf("[\n%s\n]", postElements)
}
