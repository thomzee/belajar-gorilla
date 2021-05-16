package web

import (
	"assignment3/controllers"
	"assignment3/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChatController struct {
	Base controllers.Controller
}

func (c *ChatController) Index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "chat_index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func (c *ChatController) Websocket(ctx echo.Context) error {
	models.ServeWs(c.Base.Hub, ctx)
	return nil
}