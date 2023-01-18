package comment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Content    string
	CreatedAt  time.Time
	Komentator string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
}
type CommentService interface {
	Add(newComment Core, PostID uint, token interface{}) (Core, error)
}

type CommentData interface {
	Add(newComment Core, PostID uint, UserId uint) (Core, error)
}
