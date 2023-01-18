package services

import (
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
			ID:      uint(1),
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

		// x := multipart.FileHeader{Filename: "patrick.jpg"} //data foto

		repo.On("Add", inputData).Return(resData, nil)
		srv := New(repo)
		res, err := srv.Add(inputData, pToken, nil)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Content, res.Content)
		repo.AssertExpectations(t)

	})
}
