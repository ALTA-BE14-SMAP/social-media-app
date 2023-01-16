package content

import (
	"github.com/labstack/echo/v4"
)

type CoreContent struct {
	ID      uint
	Content string
	Image   string
	Users   CoreUser
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
	Add(newContent CoreContent, token interface{}) (CoreContent, error)
	GetAll() ([]CoreContent, error)
	GetById(token interface{}) ([]CoreContent, error)
	Update(token interface{}, id uint, updatedData CoreContent) (CoreContent, error)
	Delete(token interface{}, contentId uint) error
}

type ContentData interface {
	Add(newContent CoreContent, id uint) (CoreContent, error)
	GetAll() ([]CoreContent, error)
	GetById(userId uint) ([]CoreContent, error)
	Update(userId uint, contentId uint, updatedData CoreContent) (CoreContent, error)
	Delete(userId uint, contentId uint) error
}
