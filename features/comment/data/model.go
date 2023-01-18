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
		CreatedAt:  data.CreatedAt,
		Komentator: data.User.Name,
	}
}

func CoreToData(data comment.Core) Comment {
	return Comment{
		Model: gorm.Model{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
		},
		Content:    data.Content,
		Komentator: data.Komentator,
	}
}
