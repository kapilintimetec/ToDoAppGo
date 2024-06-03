package api

import (
	"ToDoAppGo/app/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitTodoRoutes(e *echo.Echo) {
	todoGroup := e.Group("/todo")

	todoGroup.GET("", handlers.GetTodoList)
	todoGroup.POST("", handlers.CreateTodoItem)
	todoGroup.GET("/:id", handlers.GetToDoItem)
	todoGroup.PATCH("/:id", handlers.UpdateTodoItem)
	todoGroup.DELETE("/:id", handlers.DeleteTodoItem)
}
