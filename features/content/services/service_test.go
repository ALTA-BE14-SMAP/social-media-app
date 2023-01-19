package services

import (
	"errors"
	"social-media-app/features/content"
	"social-media-app/helper"
	"social-media-app/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("posting content berhasil", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      1,
			Content: "deva",
			Image:   "www.google.com",
		}
		resData := content.CoreContent{
			ID:      1,
			Content: "deva",
			Image:   "www.google.com",
		}

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		// x := multipart.FileHeader{Filename: "patrick.jpg"} //data foto

		repo.On("Add", inputData, uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Add(inputData, pToken, nil)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Content, res.Content)
		repo.AssertExpectations(t)
		// type SampleUsers struct {
		// 	ID             uint
		// 	Content        string
		// 	Image          string
		// 	UserID         uint
		// 	JumlahKomentar string
		// 	Pemilik        string
		// 	Pembuatan      string
		// }
		// inputData := SampleUsers{
		// 	ID:      1,
		// 	Content: "deva",
		// 	Image:   "www.google.com",
		// }

		// input := content.CoreContent{
		// 	Content: inputData.Content,
		// 	Image:   inputData.Image,
		// }

		// resData := content.CoreContent{
		// 	ID:      1,
		// 	Content: "deva",
		// 	Image:   "www.google.com",
		// }

		// _, token := helper.GenerateJWT(1)
		// pToken := token.(*jwt.Token)
		// pToken.Valid = true

		// // x := multipart.FileHeader{Filename: "patrick.jpg"} //data foto

		// repo.On("Add", input).Return(resData, nil).Once()
		// srv := New(repo)
		// res, err := srv.Add(input, pToken, nil)
		// assert.Nil(t, err)
		// assert.Equal(t, resData.ID, res.ID)
		// assert.Equal(t, resData.Content, res.Content)
		// repo.AssertExpectations(t)

		// input := content.CoreContent{ID: 1, Content: "test1", Image: "testing lagi"}
		// addData := content.CoreContent{ID: uint(1), Content: "test1", Image: "testing lagi"}
		// repo.On("add", uint(1), input).Return(addData, nil).Once()

		// service := New(repo)
		// _, token := helper.GenerateJWT(1)
		// pToken := token.(*jwt.Token)
		// pToken.Valid = true
		// res, err := service.Add(input, pToken, nil)
		// assert.NoError(t, err)
		// assert.Equal(t, addData.ID, res.ID)
		// assert.Equal(t, input.Content, res.Content)
		// assert.Equal(t, input.Image, res.Image)
		// repo.AssertExpectations(t)

		// repo.On("add", mock.Anything).Return(content.CoreContent{ID: uint(1), Content: "test1", Image: "testing lagi"}, nil).Once()

		// srv := New(repo)
		// input := content.CoreContent{ID: uint(1), Content: "test1", Image: "testing lagi"}
		// _, token := helper.GenerateJWT(1)
		// pToken := token.(*jwt.Token)
		// pToken.Valid = true
		// res, err := srv.Add(input, pToken, nil)
		// assert.Nil(t, err)
		// assert.NotEmpty(t, res)
		// repo.AssertExpectations(t)

	})
	t.Run("id tidak ditemukan", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      1,
			Content: "deva",
			Image:   "www.google.com",
		}
		repo.On("Add", inputData, uint(1)).Return(content.CoreContent{}, errors.New("not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "id tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("data tidak bisa diolah", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      1,
			Content: "deva",
			Image:   "www.google.com",
		}
		repo.On("Add", inputData, uint(1)).Return(content.CoreContent{}, errors.New("")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak bisa diolah")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

// func TestGetAll(t *testing.T) {
// 	repo := mocks.NewContentData(t)
// 	t.Run("oashdoiashdoiasdoi", func(t *testing.T) {
// 		resData := []content.CoreContent{
// 			{
// 				ID:      1,
// 				Content: "sedang ada di paris",
// 				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
// 			},
// 			{
// 				ID:      2,
// 				Content: "sedang ada di paris",
// 				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
// 			},
// 			{
// 				ID:      3,
// 				Content: "sedang ada di paris",
// 				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
// 			},
// 			{
// 				ID:      4,
// 				Content: "sedang ada di paris",
// 				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
// 			},
// 			{
// 				ID:      5,
// 				Content: "sedang ada di paris",
// 				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
// 			},
// 		}
// 	})

// }
