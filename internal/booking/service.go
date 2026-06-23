package booking

import (
	"go-tickets/internal/booking/dto"
	"go-tickets/internal/event"

	"github.com/google/uuid"
)

type service struct {
	bookingRepo Respository
	eventRepo   event.Repository
}

func NewService(bookingRepo Respository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func generateBookingCode() string {
	return "GT-" + uuid.New().String()
}

func (s *service) CreateBooking(userId uint, req dto.CreateRequest) (*dto.Response, error) {
	event, err := s.eventRepo.GetByID(req.EventID)
	if err != nil {
		return nil, err
	}

	if event.AvailableTickets < req.Quantity {
		return nil, ErrNotEnoughTickets
	}

	booking := &Booking{
		UserID:      userId,
		EventID:     req.EventID,
		Quantity:    req.Quantity,
		Status:      BookingConfirmed,
		TotalPrice:  req.Quantity * event.Price,
		BookingCode: generateBookingCode(),
	}

	if err := s.bookingRepo.Create(booking); err != nil {
		return nil, err
	}

	event.AvailableTickets -= req.Quantity
	if err := s.eventRepo.Update(event); err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}

func (s *service) GetMyBookings(userId uint) ([]*dto.Response, error) {
	bookings, err := s.bookingRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.Response, len(bookings))
	for i, b := range bookings {
		response[i] = b.ToResponse()
	}

	return response, nil
}
