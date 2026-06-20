package event

import (
	"errors"
	"go-tickets/internal/event/dto"
	"go-tickets/internal/httpresponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func eventErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrEventNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Event not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateEvent(c *echo.Context) error {
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

	response, err := h.service.CreateEvent(req)
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) GetEvents(c *echo.Context) error {
	events, err := h.service.GetEvents()
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, events)
}

func (h *handler) GetEventByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid event ID",
			Details: err.Error(),
		})
	}
	event, err := h.service.GetEventByID(uint(id))
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, event)
}

func (h *handler) UpdateEvent(c *echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid event ID",
			Details: err.Error(),
		})
	}
	var req dto.UpdateRequest

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

	response, err := h.service.UpdateEvent(uint(eventID), req)
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, response)
}
