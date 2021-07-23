package controller

func GetUsersController(c echo.Context) {
	users, err := database.GetUsers()
}