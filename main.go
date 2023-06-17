package main

import (
	"log"

	"github.com/abhint/linkee/database"
	"github.com/abhint/linkee/middleware"
	"github.com/abhint/linkee/router"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.NewDatabase("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	app := echo.New()

	app.Use(middleware.DatabaseMW(db))
	app.GET("/", router.IndexHandler)
	app.POST("/api", router.APIRequestHandler)
	app.GET("/:key", router.KeyRequestHandler)
	app.Start(":5000")
}
