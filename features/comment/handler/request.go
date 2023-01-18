package handler

import "social-media-app/features/comment"

type AddCommentRequest struct {
	Content string `json:"text" form:"text"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentRequest:
		cnv := data.(AddCommentRequest)
		res.Content = cnv.Content
	default:
		return nil
	}

	return &res
}
