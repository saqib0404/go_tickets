package booking

import (
	"go-tickets/internal/auth"
	"go-tickets/internal/config"
	"go-tickets/internal/event"
	"go-tickets/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	bookingRepo := NewRepository(db)
	eventRepo := event.NewRepository(db)

	service := NewService(bookingRepo, eventRepo)
	handler := NewHandler(service)

	jwtService := auth.NewJWTService(cfg.JWTSecret)

	api := e.Group("/api/v1/bookings", middleware.AuthMiddleware(jwtService))

	api.POST("", handler.CreateBooking)
	api.GET("/me", handler.GetMyBookings)
}
