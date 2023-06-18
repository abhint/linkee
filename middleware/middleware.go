package middleware

import (
	"github.com/abhint/linkee/database"
	"github.com/labstack/echo/v4"
)

func DatabaseMW(db *database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
