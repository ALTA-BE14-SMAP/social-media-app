package data

import (
	"errors"
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

func (cq *CommentQuery) Add(newComment comment.Core, PostID uint, userID uint) (comment.Core, error) {
	cnv := CoreToData(newComment)
	cnv.ContentID = PostID
	cnv.UserID = userID
	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("add comment query error :", err.Error())
		return comment.Core{}, err
	}
	newComment.ID = cnv.ID
	newComment.CreatedAt = cnv.CreatedAt.String()
	return newComment, nil
}

func (cq *CommentQuery) ListComments(PostID uint) ([]comment.Core, error) {
	res := []Comment{}

	err := cq.db.Raw(`
	SELECT c.id ,c.content, c.created_at, u.name "Komentator"
	FROM comments c 
	JOIN users u ON u.id = c.user_id
	JOIN contents c2 ON c2.id = c.content_id 
	WHERE c.deleted_at IS NULL 
	AND c2.id = ?;
	`, PostID).Scan(&res).Error
	if err != nil {
		log.Println("list book query error :", err.Error())
		return []comment.Core{}, err
	}

	return ToCoreArr(res), nil
}

func (cq *CommentQuery) Delete(commentID uint, userID uint) error {
	comment := Comment{
		Model: gorm.Model{ID: commentID},
	}
	qry := cq.db.Where("user_id = ?", userID).Delete(&comment)
	if qry.RowsAffected <= 0 {
		log.Println("delete comment query error : data not found")
		return errors.New("not found")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete comment query error :", err.Error())
		return err
	}
	return nil
}

func (cq *CommentQuery) Update(newComment comment.Core, commentID uint, userID uint) (comment.Core, error) {
	cnv := CoreToData(newComment)
	cnv.ID = commentID
	cnv.UserID = userID
	qry := cq.db.Model(&cnv).Where("user_id = ?", userID).Updates(cnv)
	log.Println(qry.RowsAffected)
	if qry.RowsAffected <= 0 {
		log.Println("update comment query error : data not found")
		return comment.Core{}, errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update comment query error :", err.Error())
		return comment.Core{}, err
	}
	// cnv.Pemilik = user.Nama
	return ToCore(cnv), nil
}
