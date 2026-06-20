package booking

import (
	"go-tickets/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

func getCurrentuserID(c *echo.Context) (uint, bool) {
	userID, ok := c.Get("user_id").(uint)
	return userID, ok
}

func bookingErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err,ErrBookingNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Booking not found",
		})
	}

	if errors.Is(err, ErrEventNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Event not found",
		})
	}

	if errors.Is(err, ErrNotEnoughTickets) {
		return c.JSON(http.StatusConflict, httpresponse.ErrorResponse{
			Code:    http.StatusConflict,
			Message: "Not enough tickets available",
		})
	}

	if errors.Is(err, ErrBookingAlreadyCancelled) {
		return c.JSON(http.StatusConflict, httpresponse.ErrorResponse{
			Code:    http.StatusConflict,
			Message: "Booking is already cancelled",
		})
	}
