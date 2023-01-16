package helper

import (
	"errors"
	"log"
	"social-media-app/features/user"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterRequest struct {
	Name     string `validate:"required" `
	Username string `validate:"required,alphanum"`
	Email    string `validate:"required,email"`
	Password string
}

type LoginUsernameRequest struct {
	Username string `validate:"required,alphanum"`
	Password string
}
type LoginEmailRequest struct {
	Email    string `validate:"required,email"`
	Password string
}

func ToRegister(data user.Core) RegisterRequest {
	return RegisterRequest{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}

func ToEmailLogin(data user.Core) LoginEmailRequest {
	return LoginEmailRequest{
		Email:    data.Email,
		Password: data.Password,
	}
}
func ToUsernameLogin(data user.Core) LoginUsernameRequest {
	return LoginUsernameRequest{
		Password: data.Password,
		Username: data.Username,
	}
}

func Validasi(data interface{}) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		msg := ""
		if strings.Contains(err.Error(), "required") {
			msg = "field required wajib diisi"
		} else if strings.Contains(err.Error(), "email") {
			msg = "format email salah"
		} else if strings.Contains(err.Error(), "Username") {
			msg = "format username salah"
		} else if strings.Contains(err.Error(), "PhoneNumber") {
			msg = "format phone_number salah"
		}
		return errors.New(msg)
	}
	return nil
}
