package posts

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createPostsOnDB(title string, content string, authorID string) error {
	connStr := "postgres://postgres:420270@localhost/blog?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)", title, content, authorID)

	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}

	return nil
}

func getAllPosts() ([]Post, error) {
	connStr := "postgres://postgres:420270@localhost/blog?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for rows.Next() {
		var post Post
		var deleted_at sql.NullTime
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &deleted_at)
		if err != nil {
			return nil, err
		}

		if deleted_at.Valid {
			post.DeletedAt = &deleted_at.Time
		}

		posts = append(posts, post)
	}

	return posts, nil
}
