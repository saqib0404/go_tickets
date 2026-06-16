package dto

import "time"

type CreateRequest struct {
	Title        string    `json:"title" validation:"required,min=2,max=150"`
	Description  string    `json:"description" validation:"omitempty,max=1000"`
	Location     string    `json:"location" validation:"required"`
	StartsAt     time.Time `json:"starts_at" validation:"required"`
	TotalTickets int       `json:"total_tickets" validation:"required,gt=0"`
	Price        int       `json:"price" validation:"required,gt=0"`
}

type UpdateRequest struct {
	Title       string    `json:"title" validation:"omitempty,min=2,max=150"`
	Description string    `json:"description" validation:"omitempty,max=1000"`
	Location    string    `json:"location" validation:"omitempty"`
	StartsAt    time.Time `json:"starts_at" validation:"omitempty"`
	Price       int       `json:"price" validation:"omitempty,gt=0"`
}
