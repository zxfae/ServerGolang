package server

import (
	"context"
	"fmt"
	"net/http"
	"real-time-backend/backend/database"
	"time"

	"github.com/gorilla/mux"
)

func LoadServer() {
	router := mux.NewRouter()

	router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", http.FileServer(http.Dir("frontend"))))

	router.HandleFunc("/", HomeHandler).Methods("GET")

	/*router.HandleFunc("/api/comments", commentsHandler).Methods("GET")
	router.HandleFunc("/api/likescomments", LikesCommentsHandler).Methods("GET")
	router.HandleFunc("/api/postscategories", PostscategoriesHandler).Methods("GET")
	router.HandleFunc("/api/postslikes", PostsLikesHandler).Methods("GET")*/

	router.HandleFunc("/api/users", userHandler).Methods("GET")

	//Client interactions
	router.HandleFunc("/api/categories", categorieHandler).Methods("GET")
	router.HandleFunc("/api/posts", postHandler).Methods("GET")
	router.HandleFunc("/api/posts/category/{categoryName}", postsByCategoryHandler).Methods("GET")
	router.HandleFunc("/api/posts/posts/{postName}", postsByPostsHandler).Methods("GET")
	//Auth session
	router.HandleFunc("/api/checkAuth", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_session_")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, valid := getSession(cookie.Value)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		if r.Method == http.MethodPost {
			nickname := r.FormValue("nickname")
			age := r.FormValue("age")
			gender := r.FormValue("gender")
			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			email := r.FormValue("email")
			password := r.FormValue("password")

			err := database.RegisterUser(ctx, nickname, age, gender, firstname, lastname, email, password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			sessionID := createSession(nickname)
			http.SetCookie(w, &http.Cookie{
				Name:     "_session_",
				Value:    sessionID,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			})
			fmt.Fprintf(w, "User %s successfully registered", nickname)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods("POST")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if r.Method == http.MethodPost {
			nickname := r.FormValue("nickname")
			password := r.FormValue("password")

			err := database.LoginUser(ctx, nickname, password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			invalidateSession(nickname)

			sessionID := createSession(nickname)
			http.SetCookie(w, &http.Cookie{
				Name:     "_session_",
				Value:    sessionID,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			})

			fmt.Fprintf(w, "User %s successfully logged in", nickname)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods("POST")

	router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_session_")
		if err != nil {
			http.Error(w, "No session found", http.StatusUnauthorized)
			return
		}

		deleteSession(cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     "_session_",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})

		fmt.Fprintf(w, "User successfully logged out")
	}).Methods("POST")
	//Test
	router.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if r.Method == http.MethodPost {
			category := r.FormValue("category")
			title := r.FormValue("title")
			description := r.FormValue("description")

			cookie, err := r.Cookie("_session_")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userSession, valid := getSession(cookie.Value)
			if !valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			username := userSession.Username

			err = database.CreatePost(ctx, username, title, description, category)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Post %s successfully created", title)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods("POST")

	serverConfig := ServerParameters(router, 10)

	fmt.Println("Server started at 127.0.0.1:8080")
	if err := serverConfig.ListenAndServe(); err != nil {
		fmt.Println("Error to serve:", err)
	}
}
