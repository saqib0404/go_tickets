package booking

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrBookingNotFound         = errors.New("booking Not found")
	ErrNotEnoughTickets        = errors.New("Not enough Tickets Available")
	ErrBookingAlreadyCancelled = errors.New("Booking already cancelled")
	ErrForbiddenBookingAccess  = errors.New(" you do not own this booking")
	ErrEventNotFound           = errors.New(" Event is not found")
)

type Respository interface {
	Create(booking *Booking) error
	GetByID(bookingId uint) (*Booking, error)
	GetByUserID(userId uint) ([]*Booking, error)
	Update(booking *Booking) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Respository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}

func (r *repository) GetByID(bookingId uint) (*Booking, error) {
	var booking Booking

	err := r.db.First(&booking, bookingId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookingNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (r *repository) GetByUserID(userId uint) ([]*Booking, error) {
	var bookings []*Booking

	err := r.db.Where("user_id = ?", userId).Find(&bookings).Error
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *repository) Update(booking *Booking) error {
	return r.db.Save(booking).Error
}
