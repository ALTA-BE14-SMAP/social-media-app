package main

import (
	"log"
	"social-media-app/config"
	"social-media-app/features/user/data"
	"social-media-app/features/user/handler"
	"social-media-app/features/user/services"

	dt "social-media-app/features/content/data"
	hd "social-media-app/features/content/handler"
	sc "social-media-app/features/content/services"

	commentData "social-media-app/features/comment/data"
	commentHandler "social-media-app/features/comment/handler"
	commentService "social-media-app/features/comment/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	contentData := dt.New(db)
	contentSrv := sc.New(contentData)
	contentHdl := hd.New(contentSrv)

	commentData := commentData.New(db)
	commentSrv := commentService.New(commentData)
	commentHdl := commentHandler.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/users", userHdl.ListUsers(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/myprofile", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Deactive(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/contents", contentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/contents", contentHdl.GetAll())
	e.GET("/contents/:id", contentHdl.GetById(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/contents/:id", contentHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/contents/:id", contentHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/comments/:idPost", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/comments/:idPost", commentHdl.ListComments())
	e.DELETE("/comments/:idPost/:idComment", commentHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
