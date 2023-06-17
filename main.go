package main

import (
	"io"
	"log"
	"text/template"

	"github.com/abhint/linkee/database"
	"github.com/abhint/linkee/middleware"
	"github.com/abhint/linkee/router"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	db, err := database.NewDatabase("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	app := echo.New()
	app.Static("/static", "public/static")
	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./public/templates/*.html")),
	}
	app.Use(middleware.DatabaseMW(db))
	app.GET("/", router.IndexHandler)
	app.POST("/api", router.APIRequestHandler)
	app.GET("/:key", router.KeyRequestHandler)
	app.Start(":5000")
}
