package event

import "go-tickets/internal/event/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateEvent(req dto.CreateRequest) (*dto.Response, error) {
	event := Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartsAt:         req.StartsAt,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}

	err := s.repo.Create(&event)
	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) GetEvents() ([]*dto.Response, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []*dto.Response
	for _, event := range events {
		responses = append(responses, event.ToResponse())
	}
	return responses, nil
}

func (s *service) GetEventByID(id uint) (*dto.Response, error) {
	event, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}
