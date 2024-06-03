package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTodoList(e echo.Context) error {
	return e.String(http.StatusOK, "ToDo list")
}

func GetToDoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo item")
}

func CreateTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo create item")
}

func UpdateTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo update item")
}

func DeleteTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo delete item")
}
