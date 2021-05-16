package routes

import (
	"assignment3/controllers"
	"assignment3/controllers/web"
	"assignment3/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func WebRoute(e *echo.Echo, db *gorm.DB) {
	// Initialize base controller
	hub := models.NewHub()
	go hub.Run()
	base := controllers.Controller{Db: db, Hub: hub}

	chatController := web.ChatController{Base: base}
	e.GET("/", chatController.Index)
	e.GET("/ws", chatController.Websocket)
}
