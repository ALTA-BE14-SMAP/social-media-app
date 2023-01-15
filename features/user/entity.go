package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Name       string `validate:"required"`
	Username   string `validate:"required" json:"username" form:"username"`
	Email      string `validate:"required,email"`
	Photo      string
	DateOfBith string
	Password   string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	// Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(newUser Core) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(newUser Core, token interface{}, image *multipart.FileHeader) (Core, error)
	// Deactive(id uint) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(newUser Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	// Deactive(id uint) error
}
