package services

import (
	"errors"
	"log"
	"social-media-app/features/content"
	"social-media-app/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type contentUseCase struct {
	qry content.ContentData
	vld *validator.Validate
}

func New2(cd content.ContentData) content.ContentService {
	return &contentUseCase{
		qry: cd,
		vld: validator.New(),
	}
}

func (cuu *contentUseCase) Add(newContent content.CoreContent, token interface{}) (content.CoreContent, error) {
	id := helper.ExtractToken(token)

	err := cuu.vld.Struct(newContent)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return content.CoreContent{}, errors.New("format input user tidak sesuai dengan arahan")
	}
	res, err := cuu.qry.Add(newContent, uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "id tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
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
			msg = "terjadi kesalahan pada server"
		}
		return []content.CoreContent{}, errors.New(msg)
	}
	return res, nil
}

func (cuu *contentUseCase) GetById(token interface{}) ([]content.CoreContent, error) {
	id2 := helper.ExtractToken(token)
	res, err := cuu.qry.GetById(uint(id2))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return []content.CoreContent{}, errors.New(msg)
	}

	return res, nil
}

func (cuu *contentUseCase) Update(token interface{}, id uint, tmp content.CoreContent) (content.CoreContent, error) {
	id2 := helper.ExtractToken(token)
	res, err := cuu.qry.Update(uint(id2), id, tmp)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
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
			msg = "terdapat masalah pada server"
		}
		return errors.New(msg)
	}

	return nil
}
