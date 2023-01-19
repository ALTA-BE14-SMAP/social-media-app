package handler

import (
	"social-media-app/features/content"
)

type ContentResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	UserID    uint   `json:"userid"`
	Name      string `json:"name"`
	Pembuatan string `json:"created_at"`
}

func ToResponse(data content.CoreContent) ContentResponse {
	return ContentResponse{
		ID:      data.ID,
		Content: data.Content,
		Image:   data.Image,
		// Pembuatan: data.Pembuatan,
	}
}

func ToResponseArr(data []content.CoreContent) []ContentResponse {
	res := []ContentResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
