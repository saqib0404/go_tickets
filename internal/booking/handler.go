package booking

import (
	"errors"
	"go-tickets/internal/booking/dto"
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
	if errors.Is(err, ErrBookingNotFound) {
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

	return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateBooking(c *echo.Context) error {
	userId, ok := getCurrentuserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var req dto.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}
	response, err := h.service.CreateBooking(userId, req)

	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response)

}

func (h *handler) GetMyBookings(c *echo.Context) error {
	userId, ok := getCurrentuserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	bookings, err := h.service.GetMyBookings(userId)
	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, bookings)
}
