package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	userhandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userhandler.CreateUser) //api/v1/auth/register
	api.POST("/login", userhandler.LoginUser)     //api/v1/auth/login
}
