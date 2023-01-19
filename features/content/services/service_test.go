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
		repo.On("Add", inputData, uint(1)).Return(content.CoreContent{}, errors.New("sembarang")).Once()
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

func TestGetAll(t *testing.T) {
	repo := mocks.NewContentData(t)
	t.Run("oashdoiashdoiasdoi", func(t *testing.T) {
		resData := []content.CoreContent{
			{
				ID:      1,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      2,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      3,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      4,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      5,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
		}
		repo.On("GetAll").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res[0].ID, resData[0].ID)
		assert.Equal(t, res[0].Content, resData[0].Content)
		assert.Equal(t, res[0].Image, resData[0].Image)
		repo.AssertExpectations(t)

	})
	t.Run("content tidak ditemukan", func(t *testing.T) {
		repo.On("GetAll").Return([]content.CoreContent{}, errors.New("not found")).Once()
		srv := New(repo)

		res, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "content tidak ditemukan")
		assert.Equal(t, res, []content.CoreContent{})
		repo.AssertExpectations(t)
	})
	t.Run("data tidak bisa diolah", func(t *testing.T) {
		repo.On("GetAll").Return([]content.CoreContent{}, errors.New("sembarang")).Once()
		srv := New(repo)

		res, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak bisa diolah")
		assert.Equal(t, res, []content.CoreContent{})
		repo.AssertExpectations(t)
	})

}
func TestGetByID(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("berhasil", func(t *testing.T) {
		resData := []content.CoreContent{
			{
				ID:      1,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      2,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      3,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      4,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:      5,
				Content: "sedang ada di paris",
				Image:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
		}
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("GetById", uint(1), uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetById(pToken, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, res[1].ID, resData[1].ID)
		assert.Equal(t, res[1].Content, resData[1].Content)
		assert.Equal(t, res[1].Image, resData[1].Image)
		repo.AssertExpectations(t)
	})
	t.Run("content tidak ditemukan", func(t *testing.T) {
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("GetById", uint(1), uint(1)).Return([]content.CoreContent{}, errors.New("not found")).Once()
		srv := New(repo)
		res, err := srv.GetById(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "content tidak ditemukan")
		assert.Equal(t, res, []content.CoreContent{})
		repo.AssertExpectations(t)
	})
	t.Run("data tidak bisa diolah", func(t *testing.T) {
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("GetById", uint(1), uint(1)).Return([]content.CoreContent{}, errors.New("sembarang")).Once()
		srv := New(repo)

		res, err := srv.GetById(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak bisa diolah")
		assert.Equal(t, res, []content.CoreContent{})
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("berhasil", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      uint(1),
			Content: "sedang di belanda",
			Image:   "www.google.com",
		}
		resData := content.CoreContent{
			ID:      uint(1),
			Content: "sedang di belanda",
			Image:   "www.google.com",
		}
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Update", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Update(pToken, uint(1), inputData, nil)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, resData.ID)
		assert.Equal(t, res.Content, resData.Content)
		assert.Equal(t, res.Image, resData.Image)
		repo.AssertExpectations(t)
	})
	t.Run("data tidak ditemukan", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      uint(1),
			Content: "sedang di belanda",
			Image:   "www.google.com",
		}
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Update", uint(1), uint(1), inputData).Return(content.CoreContent{}, errors.New("not found")).Once()
		srv := New(repo)
		res, err := srv.Update(pToken, uint(1), inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, content.CoreContent{})
		repo.AssertExpectations(t)
	})
	t.Run("data tidak bisa diolah", func(t *testing.T) {
		inputData := content.CoreContent{
			ID:      uint(1),
			Content: "sedang di belanda",
			Image:   "www.google.com",
		}
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Update", uint(1), uint(1), inputData).Return(content.CoreContent{}, errors.New("sembarang")).Once()
		srv := New(repo)
		res, err := srv.Update(pToken, uint(1), inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak bisa diolah")
		assert.Equal(t, res, content.CoreContent{})
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("berhasil", func(t *testing.T) {
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(pToken, uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})
	t.Run("data tidak ditemukan", func(t *testing.T) {
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Delete", uint(1), uint(1)).Return(errors.New("not found")).Once()
		srv := New(repo)
		err := srv.Delete(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		repo.AssertExpectations(t)

	})
	t.Run("data tidak bisa diolah", func(t *testing.T) {
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		repo.On("Delete", uint(1), uint(1)).Return(errors.New("sembarang")).Once()
		srv := New(repo)
		err := srv.Delete(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak bisa diolah")
		repo.AssertExpectations(t)

	})

}
