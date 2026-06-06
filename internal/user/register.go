package user

import (
	"go-tickets/internal/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userRepository := NewRepository(db)
	jwtService := auth.NewJWTService("")
	userService := NewService(userRepository, jwtService)
	userhandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userhandler.CreateUser) //api/v1/auth/register
	api.POST("/login", userhandler.LoginUser)     //api/v1/auth/login
}
