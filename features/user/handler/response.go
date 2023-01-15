package handler

import "social-media-app/features/user"

type UserReponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Photo      string `json:"photo"`
	DateOfBith string `json:"date_of_birth"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Username:   data.Username,
		Photo:      data.Photo,
		DateOfBith: data.DateOfBith,
	}
}
