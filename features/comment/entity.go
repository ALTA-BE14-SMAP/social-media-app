package comment

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Content    string
	CreatedAt  string
	Komentator string
	Photo      string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	ListComments() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}
type CommentService interface {
	Add(newComment Core, PostID uint, token interface{}) (Core, error)
	ListComments(PostID uint) ([]Core, error)
	Delete(commentID uint, token interface{}) error
	Update(newComment Core, commentID uint, token interface{}) (Core, error)

	// Delete(commentID uint, PostID uint, token interface{}) error
}

type CommentData interface {
	Add(newComment Core, PostID uint, UserId uint) (Core, error)
	ListComments(PostID uint) ([]Core, error)
	Delete(commentID uint, userID uint) error
	Update(newComment Core, commentID uint, userID uint) (Core, error)

	// Delete(commentID uint, PostID uint, userID uint) error
}
