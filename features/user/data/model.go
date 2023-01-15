package data

import (
	"social-media-app/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	// Books    []data.Books `gorm:"foreignKey:UserID"`
	Name       string
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	DateOfBith string
	Photo      string
	Password   string
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Username:   data.Username,
		Photo:      data.Photo,
		DateOfBith: data.DateOfBith,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:      gorm.Model{ID: data.ID},
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Username:   data.Username,
		Photo:      data.Photo,
		DateOfBith: data.DateOfBith,
	}
}
