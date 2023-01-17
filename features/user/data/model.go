package data

import (
	"social-media-app/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	// Books    []data.Books `gorm:"foreignKey:UserID"`
	Name        string
	Username    string `gorm:"unique,size:191"`
	Email       string `gorm:"unique,size:191"`
	DateOfBith  string
	Photo       string
	PhoneNumber string
	AboutMe     string
	Password    string
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		Username:    data.Username,
		Photo:       data.Photo,
		DateOfBith:  data.DateOfBith,
		PhoneNumber: data.PhoneNumber,
		AboutMe:     data.AboutMe,
	}
}

func ToCoreArr(data []Users) []user.Core {
	arrRes := []user.Core{}
	for _, v := range data {
		tmp := ToCore(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:       gorm.Model{ID: data.ID},
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		Username:    data.Username,
		Photo:       data.Photo,
		DateOfBith:  data.DateOfBith,
		PhoneNumber: data.PhoneNumber,
		AboutMe:     data.AboutMe,
	}
}
