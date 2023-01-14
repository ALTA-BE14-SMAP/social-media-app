package services

import (
	"errors"
	"log"
	"social-media-app/features/user"
	"social-media-app/helper"
	"strings"
)

type userUseCase struct {
	qry user.UserData
	// vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		// vld: validator.New(),
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	newUser.Password = helper.HashPassword(newUser.Password)
	res, err := uuc.qry.Register(newUser)
	if err != nil {

		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			log.Println("Email sudah terdaftar")
			msg = "Email sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}
