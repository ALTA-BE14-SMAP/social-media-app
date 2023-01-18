package handler

import (
	"social-media-app/features/comment"
	"time"
)

type CommentsResponse struct {
	ID         uint      `json:"id"`
	Content    string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	Komentator string    `json:"commentator"`
}

func ToResponse(data comment.Core) CommentsResponse {
	return CommentsResponse{
		ID:         data.ID,
		Content:    data.Content,
		CreatedAt:  data.CreatedAt,
		Komentator: data.Komentator,
	}
}
