package comment

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Content    string
	CreatedAt  string
	Komentator string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	ListComments() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
type CommentService interface {
	Add(newComment Core, PostID uint, token interface{}) (Core, error)
	ListComments(PostID uint) ([]Core, error)
	Delete(commentID uint, PostID uint, token interface{}) error
	// Delete(PostID uint) error
}

type CommentData interface {
	Add(newComment Core, PostID uint, UserId uint) (Core, error)
	ListComments(PostID uint) ([]Core, error)
	Delete(commentID uint, PostID uint, userID uint) error
}
