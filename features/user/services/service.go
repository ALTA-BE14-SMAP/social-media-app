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
	err := helper.Validasi(helper.ToRegister(newUser))
	if err != nil {
		return user.Core{}, err
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
	if len(newUser.Username) > 0 {
		err = helper.Validasi(helper.ToUsernameLogin(newUser))
		if err != nil {
			return "", user.Core{}, err
		}
	} else if len(newUser.Email) > 0 {
		err = helper.Validasi(helper.ToEmailLogin(newUser))
		if err != nil {
			return "", user.Core{}, err
		}
	} else {
		return "", user.Core{}, errors.New("email/username belum terdaftar")
	}

	res, err = uuc.qry.Login(newUser)

	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows") {
			msg = "email/username belum terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}
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

	if len(updateData.Email) > 0 {
		err := helper.Validasi(helper.ToEmailLogin(updateData))
		if err != nil {
			return user.Core{}, err
		}
	}

	if len(updateData.Username) > 0 {
		err := helper.Validasi(helper.ToUsernameLogin(updateData))
		if err != nil {
			return user.Core{}, err
		}
	}

	if len(updateData.PhoneNumber) > 0 {
		err := helper.Validasi(helper.ToPhoneNumber(updateData))
		if err != nil {
			return user.Core{}, err
		}
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return user.Core{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			idx := strings.Index(err.Error(), ",")
			msg := err.Error()
			return user.Core{}, errors.New("format input file size tidak diizinkan, size melebihi" + msg[idx+1:])
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
	}
	x := helper.HashPassword(updateData.Password)
	updateData.Password = x
	res, err := uuc.qry.Update(uint(id), updateData)
	// log.Println("res update qry:", res)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else if strings.Contains(err.Error(), "Duplicate") {
			msg = "email/username sudah terdaftar"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}
func (uuc *userUseCase) ListUsers() ([]user.Core, error) {
	res, err := uuc.qry.ListUsers()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Deactive(token interface{}) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("user tidak ditemukan harap login lagi")
	}
	err := uuc.qry.Deactive(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return errors.New(msg)
	}
	return nil
}
