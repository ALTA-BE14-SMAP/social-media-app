package data

import (
	"log"
	"social-media-app/features/comment"

	"gorm.io/gorm"
)

type CommentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &CommentQuery{
		db: db,
	}
}

func (cq *CommentQuery) Add(newComment comment.Core, PostID uint, UserId uint) (comment.Core, error) {
	cnv := CoreToData(newComment)
	cnv.ContentID = PostID
	cnv.UserID = UserId
	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("add comment query error :", err.Error())
		return comment.Core{}, err
	}
	newComment.ID = cnv.ID
	newComment.Komentator = cnv.User.Name
	return newComment, nil
}
