package handler

import (
	"social-media-app/features/content"
)

type RegisterReq struct {
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
}

func ToCore(data interface{}) *content.CoreContent {
	res := content.CoreContent{}

	switch data.(type) {
	case RegisterReq:
		cnv := data.(RegisterReq)
		res.Content = cnv.Content
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
