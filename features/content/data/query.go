package data

import (
	"errors"
	"log"
	"social-media-app/features/content"
	"strings"

	"gorm.io/gorm"
)

type contentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) content.ContentData {
	return &contentQuery{
		db: db,
	}
}

func (cq *contentQuery) Add(newBook content.CoreContent, id uint) (content.CoreContent, error) {
	cnv := CoreToData(newBook)
	cnv.User.ID = id

	err := cq.db.Create(&cnv).Error

	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "data tidak bisa diolah"
		}
		return content.CoreContent{}, errors.New(msg)
	}
	// newBook.Users.ID = cnv.ID
	// newBook.Users.Name = cnv.User.Name

	return newBook, nil
}

func (cq *contentQuery) GetAll() ([]content.CoreContent, error) {

	posts := []Contents{}
	comments := []content.Comment{}
	err := cq.db.Raw(`
	SELECT c.id , c.content, c.created_at "DibuatPada" , c.image, c.user_id ,u.name "Pemilik", COUNT( c2.id) "Jumlah Komentar", u.photo  
	FROM contents c 
	JOIN users u ON u.id = c.user_id 
	LEFT JOIN comments c2 ON c2.content_id = c.id 
	WHERE c2.deleted_at IS NULL
	AND c.deleted_at IS NULL 
	GROUP BY c.id ;
	`).Scan(&posts).Error
	if err != nil {
		log.Println("list book query error :", err.Error())
		return []content.CoreContent{}, err
	}

	for i := range posts {
		err := cq.db.Raw(`
		SELECT c.id ,c.content "Text", c.created_at "CreatedAt", u.name "Commentator", c2.id "ContentID", u.photo 
		FROM comments c 
		JOIN users u ON u.id = c.user_id
		JOIN contents c2 ON c2.id = c.content_id 
		WHERE c.deleted_at IS NULL 
		AND c2.id = ?
		LIMIT 3;
		`, posts[i].ID).Scan(&comments).Error
		if err != nil {
			log.Println("list book query error :", err.Error())
		}
		posts[i].Comments = append(posts[i].Comments, comments...)
	}
	// log.Println(posts)

	// var sementara []Contents

	// if err := cq.db.Preload("User").Find(&sementara).Error; err != nil {
	// 	log.Println("Get By ID query error", err.Error())
	// 	return ToCore2(sementara), err
	// }
	// Y := ToCore2(sementara)
	return ToCoresArr(posts), nil
}

func (cq *contentQuery) GetById(idUser uint, idContent uint) ([]content.CoreContent, error) {
	var sementara []Contents

	if idContent == 0 {
		if err := cq.db.Preload("User").Where("user_id = ?", idUser).Find(&sementara).Error; err != nil {
			log.Println("Get By ID query error", err.Error())
			return ToCore2(sementara), err
		}
		X := ToCore2(sementara)
		return X, nil
	}

	if err := cq.db.Preload("User").Where("user_id = ? AND id = ?", idUser, idContent).Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCore2(sementara), err
	}
	X := ToCore2(sementara)
	return X, nil
}

func (cq *contentQuery) Update(userId uint, contentId uint, updatedData content.CoreContent) (content.CoreContent, error) {
	cnv := CoreToData(updatedData)
	qry := cq.db.Where("user_id = ? AND id = ?", userId, contentId).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update content query error : data not found")
		return content.CoreContent{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return content.CoreContent{}, err
	}
	// cnv.ID = id
	Y := cnv.ToCore()
	return Y, nil
}

func (cq *contentQuery) Delete(userId uint, contentId uint) error {
	var sementara Contents
	if err := cq.db.Where("user_id = ? AND id = ?", userId, contentId).Delete(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return err
	}
	return nil
}
