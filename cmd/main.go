package main

import (
	"fmt"
	"go-tickets/internal/config"
	"go-tickets/internal/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100); not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(225); uniqueIndex; not null"`
	Password string `json:"password" validate:"required" gorm:"type:varchar(100); not null"`
}
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func main() {
	config := config.LoadConfig()

	dsn := config.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("failed to connect to db")
	} else {
		fmt.Println("Connected to db")
	}

	db.AutoMigrate(&User{})

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userhandler := user.NewHandler(userService)

	e.POST("/users", userhandler.CreateUser)

	if err := e.Start(":5000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
