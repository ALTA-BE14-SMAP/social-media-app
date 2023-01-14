package data

import (
	"social-media-app/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	// Books    []data.Books `gorm:"foreignKey:UserID"`
	Name     string
	Email    string
	Password string
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
