package handler

import (
	"net/http"
	"social-media-app/features/user"
	"social-media-app/helper"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		// res, err := uc.srv.Register(*ToCore(input))
		_, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		// return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("berhasil mendaftar", ToResponse(res)))
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("success add data"))
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		// token, res, err := uc.srv.Login(input.Email, input.Password)
		token, _, err := uc.srv.Login(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("success login", "", token))
		// return c.JSON(http.StatusOK, helper.PrintSuccessReponse("success login", ToResponse(res), token))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read form fields
		input := user.Core{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah, cek json")
		}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}

		token := c.Get("user")
		// res, err := uc.srv.Update(input, token, file)
		_, err = uc.srv.Update(input, token, file)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		// return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil update profil", ToResponse(res)))
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil update profil"))
	}
}

func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil lihat profil", ToResponse(res)))
	}
}
func (uc *userControll) ListUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.srv.ListUsers()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil lihat get users", ToGetUsersResArr(res)))
	}
}

func (uc *userControll) Deactive() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := uc.srv.Deactive(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil menonaktifkan akun"))
	}
}
