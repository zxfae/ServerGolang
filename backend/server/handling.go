package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-backend/backend/database"
	"text/template"
)

func categorieHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := database.GetCategoriesTable()
	if err != nil {
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetPostTable()
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := database.GetCommentsTable()
	if err != nil {
		http.Error(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LikesCommentsHandler(w http.ResponseWriter, r *http.Request) {
	likesComments, err := database.GetLikesCommentsTable()
	if err != nil {
		fmt.Println("de")
		http.Error(w, "Error fetching likesComments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(likesComments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostscategoriesHandler(w http.ResponseWriter, r *http.Request) {
	PostCategories, err := database.GetPostCategoriesTable()
	if err != nil {
		http.Error(w, "Error fetching post categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(PostCategories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsLikesHandler(w http.ResponseWriter, r *http.Request) {
	PostLikes, err := database.GetPostLikesTable()
	if err != nil {
		fmt.Println("HERE")

		http.Error(w, "Error fetching postLike", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(PostLikes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUserTable()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
