package data

import (
	"errors"
	"log"
	"social-media-app/features/user"
	"strings"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Duplicate(email string) error {
	res := Users{}
	qry := uq.db.Where("email = ?", email).First(&res)
	if qry.RowsAffected <= 0 {
		// log.Println("select user query error : data not found")
		// return errors.New("record not found")
		return nil
	}
	err := qry.Error
	if err != nil {
		log.Println("select email query error :", err.Error())
		return errors.New("server error")
	}
	if len(res.Email) > 0 {
		return errors.New("email is duplicated")
	}
	return nil
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cnv := CoreToData(newUser)
	err := uq.Duplicate(newUser.Email)
	if err != nil {
		log.Println("register query error", err.Error())
		return user.Core{}, err
	}
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email is duplicated"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return user.Core{}, nil
}
