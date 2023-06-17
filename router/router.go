package router

import (
	"net/http"

	"github.com/abhint/linkee/database"
	"github.com/labstack/echo/v4"
)

type URL string

type URLStruct struct {
	URL `json:"url"`
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "app.html", nil)
}

func APIRequestHandler(c echo.Context) error {
	mapping := &database.UrlMapping{}
	db := c.Get("db").(*database.Database)
	if db == nil {
		return c.JSON(http.StatusInternalServerError, "Database connection is nil")
	}

	urlStruct := &URLStruct{}
	if err := c.Bind(urlStruct); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error binding URL")
	}

	if urlStruct.URL == "" {
		return c.JSON(http.StatusBadRequest, "URL is empty")
	}
	if err := db.UrlMappings.Insert(string(urlStruct.URL), mapping); err != nil {
		return c.JSON(http.StatusOK, "Error inserting URL")
	}
	return c.JSON(http.StatusOK, mapping)
}

func KeyRequestHandler(c echo.Context) error {
	key := c.Param("key")
	mapping := &database.UrlMapping{}
	db := c.Get("db").(*database.Database)
	if err := db.UrlMappings.Select(key, mapping); err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	url := *mapping.Url
	return c.Redirect(http.StatusMovedPermanently, url)
}
