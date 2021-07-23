package routes

import (
	"rest_api/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/users/", controller.GetUsersController)
	e.GET("/users/:id", controller.GetOneUserController)
	e.POST("/users/", controller.CreateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
}
