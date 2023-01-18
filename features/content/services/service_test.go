package services

import (
	"mime/multipart"
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
			Content: "deva",
			Image:   "www.google.com",
		}

		resData := content.CoreContent{
			ID:      uint(1),
			Content: "deva",
			Image:   "www.google.com",
		}

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		x := multipart.FileHeader{Filename: "patrick"} //data foto

		repo.On("Add", inputData).Return(resData, nil)
		srv := New2(repo)
		res, err := srv.Add(inputData, pToken, &x)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Content, res.Content)
		assert.Equal(t, resData.Image, res.Image)
		repo.AssertExpectations(t)

	})
}
