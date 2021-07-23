package routes

import (
	"rest_api/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/users/", controller.GetUserControllers)
	e.GET("/users/:id", controller.GetOneUserControllers)
	e.POST("/users/", controller.CreateUserControllers)
	e.PUT("/users/:id", controller.EditUserControllers)
	e.DELETE("/users/:id", controller.DeleteUserControllers)
}
