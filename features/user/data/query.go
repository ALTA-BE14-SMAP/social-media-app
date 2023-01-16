package data

import (
	"database/sql"
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

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}
	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Login(newUser user.Core) (user.Core, error) {

	var (
		res Users
		row *sql.Row
	)
	log.Println(newUser)
	if len(newUser.Email) > 0 {
		row = uq.db.Raw(`
		SELECT u.id, u.password 
		FROM users u 
		WHERE u.email = ?
		AND deleted_at IS NULL;
		`, newUser.Email).Row()
	} else if len(newUser.Username) > 0 {
		row = uq.db.Raw(`
		SELECT u.id, u.password 
		FROM users u 
		WHERE u.username = ?
		AND deleted_at IS NULL;
		`, newUser.Username).Row()
	} else {
		return user.Core{}, errors.New("record not found")
	}

	err := row.Scan(&res.ID, &res.Password)
	if err != nil {
		log.Println("login query error :", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}

func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := Users{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("get by id query error", err.Error())
		return user.Core{}, err
	}
	return ToCore(res), nil

}

func (uq *userQuery) Update(id uint, updateData user.Core) (user.Core, error) {
	cnv := CoreToData(updateData)
	qry := uq.db.Model(&cnv).Where("id = ?", id).Updates(cnv)
	if qry.RowsAffected <= 0 {
		return user.Core{}, errors.New("record not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update user query error :", err.Error())
		return user.Core{}, err
	}
	cnv.ID = id
	return ToCore(cnv), nil
}

func (uq *userQuery) ListUsers() ([]user.Core, error) {
	res := []Users{}
	err := uq.db.Select("id", "username", "photo").Find(&res).Error
	if err != nil {
		log.Println("Mybook query error :", err.Error())
		return []user.Core{}, err
	}
	if len(res) == 0 {
		return []user.Core{}, errors.New("record not found")
	}
	return ToCoreArr(res), nil
}
