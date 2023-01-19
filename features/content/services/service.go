package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"social-media-app/features/content"
	"social-media-app/helper"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type contentUseCase struct {
	qry content.ContentData
	vld *validator.Validate
}

func New(cd content.ContentData) content.ContentService {
	return &contentUseCase{
		qry: cd,
		vld: validator.New(),
	}
}

func (cuu *contentUseCase) Add(newContent content.CoreContent, token interface{}, file *multipart.FileHeader) (content.CoreContent, error) {
	id := helper.ExtractToken(token)
	fmt.Println("======service=====")
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return content.CoreContent{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file size tidak diizinkan")
		}
		extension, err := helper.CheckFileExtension(file.Filename)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file type tidak diizinkan")
		}
		filename := "images/profile/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

		photo, err := helper.UploadImageToS3(filename, src)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file type tidak dapat diupload")
		}

		newContent.Image = photo

		defer src.Close()
	}

	// err := cuu.vld.Struct(newContent)
	// if err != nil {
	// 	log.Println(err)
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		log.Println(err)
	// 	}
	// 	return content.CoreContent{}, errors.New("field required wajib diisi")
	// }
	res, err := cuu.qry.Add(newContent, uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "id tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return content.CoreContent{}, errors.New(msg)
	}

	return res, nil

}

func (cuu *contentUseCase) GetAll() ([]content.CoreContent, error) {
	res, err := cuu.qry.GetAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []content.CoreContent{}, errors.New(msg)
	}
	return res, nil
}

func (cuu *contentUseCase) GetById(token interface{}, tes uint) ([]content.CoreContent, error) {
	id2 := helper.ExtractToken(token)
	res, err := cuu.qry.GetById(uint(id2), tes)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []content.CoreContent{}, errors.New(msg)
	}

	return res, nil
}

func (cuu *contentUseCase) Update(token interface{}, id uint, tmp content.CoreContent, file *multipart.FileHeader) (content.CoreContent, error) {
	id2 := helper.ExtractToken(token)
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return content.CoreContent{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file size tidak diizinkan")
		}
		extension, err := helper.CheckFileExtension(file.Filename)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file type tidak diizinkan")
		}
		filename := "images/profile/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

		photo, err := helper.UploadImageToS3(filename, src)
		if err != nil {
			return content.CoreContent{}, errors.New("format input file type tidak dapat diupload")
		}

		tmp.Image = photo

		defer src.Close()
	}
	res, err := cuu.qry.Update(uint(id2), id, tmp)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return content.CoreContent{}, errors.New(msg)
	}

	return res, nil
}

func (cuu *contentUseCase) Delete(token interface{}, contentId uint) error {
	userId := helper.ExtractToken(token)
	err := cuu.qry.Delete(uint(userId), contentId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "data tidak bisa diolah"
		}
		return errors.New(msg)
	}

	return nil
}
