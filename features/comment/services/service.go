package services

import (
	"errors"
	"social-media-app/features/comment"
	"social-media-app/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type commentUseCase struct {
	qry comment.CommentData
	vld *validator.Validate
}

func New(ud comment.CommentData) comment.CommentService {
	return &commentUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (cuc *commentUseCase) Add(newComment comment.Core, PostID uint, token interface{}) (comment.Core, error) {
	err := helper.Validasi(helper.ToComment(newComment))
	if err != nil {
		return comment.Core{}, err
	}

	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comment.Core{}, errors.New("user tidak ditemukan")
	}
	res, err := cuc.qry.Add(newComment, PostID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return comment.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cuc *commentUseCase) ListComments(PostID uint) ([]comment.Core, error) {
	res, err := cuc.qry.ListComments(PostID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []comment.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cuc *commentUseCase) Delete(commentID uint, token interface{}) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user tidak ditemukan")
	}
	err := cuc.qry.Delete(commentID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "comment tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return errors.New(msg)
	}
	return nil
}
func (cuc *commentUseCase) Update(newComment comment.Core, commentID uint, token interface{}) (comment.Core, error) {
	err := helper.Validasi(helper.ToComment(newComment))
	if err != nil {
		return comment.Core{}, err
	}

	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comment.Core{}, errors.New("user tidak ditemukan")
	}

	res, err := cuc.qry.Update(newComment, commentID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "comment tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return comment.Core{}, errors.New(msg)
	}

	return res, nil
}
