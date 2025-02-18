package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/orhanozkercin/goblog/blog"
)

func main() {
	// api/v1/post-titles?limit=asfsaf GET -> get all post titles
	http.HandleFunc("GET /api/v1/post-titles", func(w http.ResponseWriter, r *http.Request) {
		var limit blog.LimitType
		limitStr := r.URL.Query().Get("limit")

		if limitStr == "" {
			limit.HasValue = false
		} else {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				http.Error(w, "Invalid limit", http.StatusBadRequest)
				return
			}
			limit.HasValue = true
			limit.Value = limitInt
		}

		titles := blog.GetBlogTitles(limit)
		err := json.NewEncoder(w).Encode(titles)

		if err != nil {
			http.Error(w, "Failed to encode titles", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
