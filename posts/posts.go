package posts

import (
	"errors"
	"time"
)

const (
	TitleMaxLength   = 255
	ContentMaxLength = 10000
)

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	AuthorID  string     `json:"authorID"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type CreatePostRequestPayload struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorID"`
}

var ErrInvalidTitleLength = errors.New("title length must be less than 255 characters")
var ErrInvalidContentLength = errors.New("content length must be less than 10000 characters")

func CreatePost(post CreatePostRequestPayload) error {
	if len(post.Title) > TitleMaxLength {
		return ErrInvalidTitleLength
	}

	if len(post.Content) > ContentMaxLength {
		return ErrInvalidContentLength
	}

	err := createPostsOnDB(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPosts() ([]Post, error) {
	return getAllPosts()
}
