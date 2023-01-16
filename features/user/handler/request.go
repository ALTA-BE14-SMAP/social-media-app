package handler

import (
	"social-media-app/features/user"
)

type LoginRequest struct {
	Username string `validate:"required,alphanum" json:"username" form:"username"`
	Email    string `validate:"required,email" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Email = cnv.Email
		res.Username = cnv.Username
		res.Name = cnv.Name
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
