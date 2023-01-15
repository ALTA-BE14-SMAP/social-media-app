package services

import (
	"errors"
	"log"
	"mime/multipart"
	"social-media-app/features/user"
	"social-media-app/helper"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	err := uuc.vld.Struct(newUser)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return user.Core{}, errors.New("format input user tidak sesuai dengan arahan")
	}

	newUser.Password = helper.HashPassword(newUser.Password)
	res, err := uuc.qry.Register(newUser)
	if err != nil {

		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email/username sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Login(newUser user.Core) (string, user.Core, error) {
	var (
		res user.Core
		err error
	)

	res, err = uuc.qry.Login(newUser)

	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "email/password belum terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}
	log.Println("ini dari db:", res.Password)
	if err := helper.ComparePassword(res.Password, newUser.Password); err != nil {
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	useToken, _ := helper.GenerateJWT(int(res.ID))

	return useToken, res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user tidak ditemukan harap login lagi")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user tidak ditemukan harap login lagi"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(updateData user.Core, token interface{}, file *multipart.FileHeader) (user.Core, error) {

	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user tidak ditemukan harap login lagi")
	}

	src, err := file.Open()
	if err != nil {
		return user.Core{}, errors.New("format input file tidak dapat dibuka")
	}
	err = helper.CheckFileSize(file.Size)
	if err != nil {
		return user.Core{}, errors.New("format input file size tidak diizinkan")
	}
	extension, err := helper.CheckFileExtension(file.Filename)
	if err != nil {
		return user.Core{}, errors.New("format input file type tidak diizinkan")
	}
	filename := "images/profile/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

	photo, err := helper.UploadImageToS3(filename, src)
	if err != nil {
		return user.Core{}, errors.New("format input file type tidak dapat diupload")
	}

	updateData.Photo = photo

	defer src.Close()

	res, err := uuc.qry.Update(uint(id), updateData)
	// log.Println("res update qry:", res)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "buku tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}
