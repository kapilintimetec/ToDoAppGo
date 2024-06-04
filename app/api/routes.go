package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is OK!!!")
	})
	InitTodoRoutes(e)
}
