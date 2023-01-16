package handler

import "social-media-app/features/user"

type UserReponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Photo       string `json:"photo"`
	DateOfBith  string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	AboutMe     string `json:"about_me"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Username:    data.Username,
		Photo:       data.Photo,
		DateOfBith:  data.DateOfBith,
		PhoneNumber: data.PhoneNumber,
		AboutMe:     data.AboutMe,
	}
}

// get users
type GetUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
}

func ToGetUserResponse(data user.Core) GetUserResponse {
	return GetUserResponse{
		ID:       data.ID,
		Username: data.Username,
		Photo:    data.Photo,
	}
}

func ToGetUsersResArr(data []user.Core) []GetUserResponse {
	res := []GetUserResponse{}
	for _, v := range data {
		tmp := ToGetUserResponse(v)
		res = append(res, tmp)
	}
	return res
}
