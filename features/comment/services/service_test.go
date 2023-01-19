package services

import (
	"errors"
	"social-media-app/features/comment"
	"social-media-app/helper"
	"social-media-app/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewCommentData(t)
	var postID uint = 1
	t.Run("success add comment", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		resData := comment.Core{
			ID:        1,
			Content:   "Sy jg pengen punya app kayak gini",
			CreatedAt: "2023-01-18 18:33:24.388 +0700 WIB",
		}
		repo.On("Add", inputData, postID, uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(inputData, postID, pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Content, res.Content)
		assert.Equal(t, resData.CreatedAt, res.CreatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		repo.On("Add", inputData, postID, uint(1)).Return(comment.Core{}, errors.New("server error")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(inputData, postID, pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("field required wajib diisi", func(t *testing.T) {
		inputData := comment.Core{
			Content: "",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(inputData, postID, pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wajib diisi")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Add(inputData, postID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		repo.On("Add", inputData, postID, uint(1)).Return(comment.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(inputData, postID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

func TestListComments(t *testing.T) {
	repo := mocks.NewCommentData(t)
	var postID uint = 1
	t.Run("Berhasil lihat get users", func(t *testing.T) {
		resData := []comment.Core{
			{
				ID:         1,
				Content:    "ini baru bisa di update",
				CreatedAt:  "2023-01-18 18:33:24.388 +0700 WIB",
				Komentator: "Budi Sukses",
			},
			{
				ID:         2,
				Content:    "Sy jg pengen punya app kayak gini",
				CreatedAt:  "2023-01-18 18:51:15.698 +0700 WIB",
				Komentator: "Budi Sukses",
			},
			{
				ID:         3,
				Content:    "Tutorial yg paling banyak dicari tentang golang menurut saya adalah cara Deploy ke production di hosting yg murah untuk produktion,",
				CreatedAt:  "2023-01-19 06:01:41.037 +0700 WIB",
				Komentator: "Budi Sukses",
			},
		}
		repo.On("ListComments", postID).Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.ListComments(postID)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, res[0].ID)
		assert.Equal(t, resData[0].Content, res[0].Content)
		assert.Equal(t, resData[0].CreatedAt, res[0].CreatedAt)
		assert.Equal(t, resData[0].Komentator, res[0].Komentator)
		repo.AssertExpectations(t)
	})
	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := []comment.Core{}
		repo.On("ListComments", postID).Return(resData, errors.New("record not found")).Once()

		srv := New(repo)
		res, err := srv.ListComments(postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []comment.Core{})
		repo.AssertExpectations(t)
	})
	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := []comment.Core{}
		repo.On("ListComments", postID).Return(resData, errors.New("query error")).Once()

		srv := New(repo)
		res, err := srv.ListComments(postID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []comment.Core{})
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewCommentData(t)
	var commentID uint = 1
	t.Run("User tidak ditemukan", func(t *testing.T) {
		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(commentID, pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("User tidak ditemukan diquery", func(t *testing.T) {
		repo.On("Delete", commentID, uint(1)).Return(errors.New("record not found")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(commentID, pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {

		repo.On("Delete", commentID, uint(1)).Return(errors.New("query error")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(commentID, pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Berhasil menonaktifkan akun", func(t *testing.T) {
		repo.On("Delete", commentID, uint(1)).Return(nil).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(commentID, pToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewCommentData(t)
	var commentID uint = 1
	t.Run("Berhasil update comment", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		resData := comment.Core{
			ID:        1,
			Content:   "Sy jg pengen punya app kayak gini",
			CreatedAt: "2023-01-18 18:33:24.388 +0700 WIB",
		}
		repo.On("Update", inputData, commentID, uint(4)).Return(resData, nil).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, commentID, pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Update(inputData, commentID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("Comment tidak ditemukan", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}

		repo.On("Update", inputData, commentID, uint(4)).Return(comment.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, commentID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := comment.Core{
			Content: "Sy jg pengen punya app kayak gini",
		}
		repo.On("Update", inputData, commentID, uint(4)).Return(comment.Core{}, errors.New("terdapat masalah pada server")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, commentID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("field required wajib diisi", func(t *testing.T) {
		inputData := comment.Core{
			Content: "",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, commentID, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wajib diisi")
		assert.Equal(t, uint(0), res.ID)
	})

}
