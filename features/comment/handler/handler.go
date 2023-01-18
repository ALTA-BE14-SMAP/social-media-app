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
		res, err := ch.srv.Add(*cnv, uint(postID), c.Get("user"))
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("success add comment", ToResponse(res)))
	}
}
