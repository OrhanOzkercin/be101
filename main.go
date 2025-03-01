package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orhanozkercin/goblog/posts"
)

func main() {
	http.HandleFunc("POST /posts", func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var post posts.CreatePostRequestPayload
		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// create post
		err = posts.CreatePost(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// return post
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	})
	http.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		// get all posts
		posts, err := posts.GetAllPosts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// return posts
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
