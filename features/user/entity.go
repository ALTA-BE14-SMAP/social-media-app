package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Name        string `validate:"required" json:"name" form:"name"`
	Username    string `validate:"required" json:"username" form:"username"`
	Email       string `validate:"required,email" json:"email" form:"email"`
	Photo       string
	DateOfBith  string `json:"date_of_birth" form:"date_of_birth"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	AboutMe     string `json:"about_me" form:"about_me"`
	Password    string `json:"password" form:"password"`
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	ListUsers() echo.HandlerFunc

	// Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(newUser Core) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(newUser Core, token interface{}, image *multipart.FileHeader) (Core, error)
	ListUsers() ([]Core, error)
	// Deactive(id uint) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(newUser Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	ListUsers() ([]Core, error)
	// Deactive(id uint) error
}
