package booking

import "go-tickets/internal/event"

type service struct {
	bookingRepo Respository
	eventRepo   event.Repository
}

func NewService(bookingRepo Respository) *service {
	return &service{
		bookingRepo: bookingRepo,
	}
}
