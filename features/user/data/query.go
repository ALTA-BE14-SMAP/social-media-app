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
	WHERE u.email = ?
	AND deleted_at IS NULL;
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
	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	log.Println(email)
	res := Users{}
	row := uq.db.Raw(`
	SELECT u.id, u.name, u.email, u.password 
	FROM users u 
	WHERE u.email = ?
	AND deleted_at IS NULL;
	`, email).Row()
	row.Scan(&res.ID, &res.Name, &res.Email, &res.Password)
	log.Println(res)

	// if res.ID <= 0 {
	// 	return user.Core{}, errors.New("record not found")
	// }
	return ToCore(res), nil
}
