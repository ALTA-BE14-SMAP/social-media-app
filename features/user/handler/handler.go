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
