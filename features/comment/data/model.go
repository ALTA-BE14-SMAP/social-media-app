package data

import (
	"social-media-app/features/comment"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content    string
	Komentator string
	UserID     uint
	ContentID  uint
	User       User     `gorm:"foreignKey:UserID;references:ID;"`
	Contents   Contents `gorm:"foreignKey:UserID;references:ID;"`
}

type User struct {
	gorm.Model
	Name    string
	Comment []Comment
}

type Contents struct {
	gorm.Model
}

func ToCore(data Comment) comment.Core {
	return comment.Core{
		ID:         data.ID,
		Content:    data.Content,
		CreatedAt:  data.CreatedAt.String(),
		Komentator: data.Komentator,
	}
}

func CoreToData(data comment.Core) Comment {
	return Comment{
		Content:    data.Content,
		Komentator: data.Komentator,
	}
}

func ToCoreArr(data []Comment) []comment.Core {
	arrRes := []comment.Core{}
	for _, v := range data {
		tmp := ToCore(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
