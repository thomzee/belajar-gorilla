package main

import (
	"assignment3/configs"
	"assignment3/routes"
	"github.com/alecthomas/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"os"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	// App environment start here ...
	configs.Env()
	appPort := os.Getenv("APP_PORT")

	// Database initialization
	configs.InitDb()
	db := configs.GetDb()

	// Echo initialization
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "resources/plugins")

	// App global middleware start here ...
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// App routes start here ...
	routes.WebRoute(e, db)

	// App runner start here ...
	e.Logger.Fatal(e.Start(":" + appPort))
}
