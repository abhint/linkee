package main

import "github.com/labstack/echo/v4"

func main() {
	app := echo.New()

	app.Start(":5000")
}
