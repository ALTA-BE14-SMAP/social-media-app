package data

import (
	"social-media-app/features/content"

	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	Content        string
	Image          string
	UserID         uint
	JumlahKomentar uint
	Pemilik        string
	DibuatPada     string
	Comments       []content.Comment `gorm:"foreignKey:ContentID;references:ID"`
	Photo          string
	User           User
}

type User struct {
	gorm.Model
	Name        string
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	DateOfBith  string
	Photo       string
	PhoneNumber string
	AboutMe     string
	Password    string
	Contentss   []Contents
}

func ToCores(data Contents) content.CoreContent {
	return content.CoreContent{
		ID:             data.ID,
		Content:        data.Content,
		Image:          data.Image,
		UserID:         data.UserID,
		JumlahKomentar: data.JumlahKomentar,
		Pemilik:        data.Pemilik,
		CreatedAt:      data.DibuatPada,
		Comments:       data.Comments,
		Photo:          data.Photo,
	}
}

func ToCoresArr(data []Contents) []content.CoreContent {
	arrRes := []content.CoreContent{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}

func (data *Contents) ToCore() content.CoreContent {
	return content.CoreContent{
		ID:             data.ID,
		Content:        data.Content,
		Image:          data.Image,
		UserID:         data.User.ID,
		JumlahKomentar: data.JumlahKomentar,
		Pemilik:        data.User.Name,
		// Pembuatan:      data.CreatedAt.String(),
		// Users: content.CoreUser{
		// 	ID:   data.User.ID,
		// 	Name: data.User.Name,
		// 	// Username: data.User.Username,
		// 	// Email: data.User.Email,
		// 	// DateOfBith: data.User.DateOfBith,
		// 	// Photo: data.User.Photo,
		// 	// PhoneNumber: data.User.PhoneNumber,
		// 	// AboutMe: data.User.AboutMe,
		// 	// Password: data.User.Password,
		// },
	}
}

func CoreToData(data content.CoreContent) Contents {
	return Contents{
		Model:   gorm.Model{ID: data.ID},
		Content: data.Content,
		Image:   data.Image,
		// UserID:  data.Users.ID,
	}
}

func ToCore2(data []Contents) []content.CoreContent {
	var tes []content.CoreContent
	for _, v := range data {
		tes = append(tes, v.ToCore())
	}
	return tes
}
