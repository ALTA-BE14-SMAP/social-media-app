package handler

import "social-media-app/features/comment"

type AddCommentRequest struct {
	Content    string `json:"text" form:"text"`
	Komentator string `json:"commentator" form:"commentator"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentRequest:
		cnv := data.(AddCommentRequest)
		res.Content = cnv.Content
		res.Komentator = cnv.Komentator
	default:
		return nil
	}

	return &res
}
