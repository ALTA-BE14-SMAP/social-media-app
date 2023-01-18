package handler

import (
	"log"
	"net/http"
	"social-media-app/features/comment"
	"social-media-app/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentHandle struct {
	srv comment.CommentService
}

func New(cs comment.CommentService) comment.CommentHandler {
	return &commentHandle{
		srv: cs,
	}
}

func (ch *commentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCommentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		id := c.Param("idPost")
		postID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("trouble convert param id post:  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)
		// res, err := ch.srv.Add(*cnv, uint(postID), c.Get("user"))
		_, err = ch.srv.Add(*cnv, uint(postID), c.Get("user"))
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("success add comment"))
		// return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("success add comment", ToResponse(res)))
	}
}

func (ch *commentHandle) ListComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("idPost")
		postID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("trouble convert param id post:  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		res, err := ch.srv.ListComments(uint(postID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("Berhasil melihat list comments", ToResponseArr(res)))
	}
}

func (ch *commentHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("idPost")
		postID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("trouble convert param id post:  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		id = c.Param("idComment")
		commentID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("trouble convert param id comment:  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		token := c.Get("user")
		err = ch.srv.Delete(uint(commentID), uint(postID), token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Berhasil delete comment"))
	}
}
