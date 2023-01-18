package handler

import (
	"social-media-app/features/comment"
)

type CommentsResponse struct {
	ID         uint   `json:"id"`
	Content    string `json:"text"`
	CreatedAt  string `json:"created_at"`
	Komentator string `json:"commentator"`
}

func ToResponse(data comment.Core) CommentsResponse {
	return CommentsResponse{
		ID:         data.ID,
		Content:    data.Content,
		CreatedAt:  data.CreatedAt,
		Komentator: data.Komentator,
	}
}

func ToResponseArr(data []comment.Core) []CommentsResponse {
	res := []CommentsResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
