package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var contents []byte

func practice() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/png", contents)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
