package handler

import (
	"fmt"
	"log"
	"net/http"
	"social-media-app/features/content"
	"social-media-app/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type contentControll struct {
	srv content.ContentService
}

func New2(srv content.ContentService) content.ContentHandler {
	return &contentControll{
		srv: srv,
	}
}

func (cc *contentControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := RegisterReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "kesalahan input")
		}
		newContent := ToCore(input)
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}
		fmt.Println("====handler======")
		_, err = cc.srv.Add(*newContent, token, file)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("posting content berhasil"))
	}
}

func (cc *contentControll) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.srv.GetAll()
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil menampilkan content", ToResponseArr(res)))
	}
}

func (cc *contentControll) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		tes, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Data not found"))
		}
		token := c.Get("user")
		res, err2 := cc.srv.GetById(token, uint(tes))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil menampilkan content", ToResponseArr(res)))
	}
}

func (cc *contentControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		tes, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Error binding data"))
		}
		token := c.Get("user")
		tmp := RegisterReq{}
		if err := c.Bind(&tmp); err != nil {
			return c.JSON(http.StatusBadRequest, "kesalahan input")
		}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}
		_, err2 := cc.srv.Update(token, uint(tes), *ToCore(tmp), file)
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil update content"))
	}
}

func (cc *contentControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tes, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Kesalahan input"))
		}
		token := c.Get("user")
		err2 := cc.srv.Delete(token, uint(tes))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil delete content"))
	}
}
