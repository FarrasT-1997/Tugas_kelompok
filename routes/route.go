package routes

import (
	"rest_api/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/users/", controller.GetUserController)
	e.GET("/users/:id", controller.GetOneUserController)
	e.POST("/users/", controller.CreateUserController)
	e.PUT("/users/:id", controller.EditUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
}
