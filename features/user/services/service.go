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
			msg = "email sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {

	res, err := uuc.qry.Login(email)

	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}
	log.Println("ini dari db:", res.Password)
	if err := helper.ComparePassword(res.Password, password); err != nil {
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	useToken, _ := helper.GenerateJWT(int(res.ID))

	return useToken, res, nil
}
