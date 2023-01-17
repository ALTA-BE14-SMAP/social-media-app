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

	contentData := dt.New2(db)
	contentSrv := sc.New2(contentData)
	contentHdl := hd.New2(contentSrv)

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
	e.GET("/contents", contentHdl.GetById(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/contents/:id", contentHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/contents/:id", contentHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
