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
	var res uint
	row := uq.db.Raw(`
	SELECT u.id 
	FROM users u 
	WHERE u.email = ?;
	`, email).Row()
	row.Scan(&res)
	if res > 0 {
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
