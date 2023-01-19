package content

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CoreContent struct {
	ID             uint   `json:"id" form:"id"`
	Content        string `validate:"required" json:"content" form:"content"`
	Image          string `json:"image" form:"image"`
	UserID         uint   `json:"user_id" form:"user_id"`
	JumlahKomentar string `json:"number_of_comments" form:"number_of_comments"`
	Pemilik        string `json:"who_post" form:"who_post"`
	// Pembuatan      string
	Comments []Comment
}

type Comment struct {
	ID          uint   `json:"id" form:"id"`
	Text        string `json:"text" form:"text"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	Commentator string `json:"comentator" form:"comentator"`
	ContentID   uint   `json:"id_post" form:"id_post"`
}

type CoreUser struct {
	ID   uint
	Name string
}

type ContentHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ContentService interface {
	Add(newContent CoreContent, token interface{}, image *multipart.FileHeader) (CoreContent, error)
	GetAll() ([]CoreContent, error)
	GetById(token interface{}, tes uint) ([]CoreContent, error)
	Update(token interface{}, id uint, updatedData CoreContent, file *multipart.FileHeader) (CoreContent, error)
	Delete(token interface{}, contentId uint) error
}

type ContentData interface {
	Add(newContent CoreContent, id uint) (CoreContent, error)
	GetAll() ([]CoreContent, error)
	GetById(userId uint, tes uint) ([]CoreContent, error)
	Update(userId uint, contentId uint, updatedData CoreContent) (CoreContent, error)
	Delete(userId uint, contentId uint) error
}
