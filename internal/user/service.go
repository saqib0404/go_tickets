package user

import "go-tickets/internal/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := user.hashPassword(req.Password); err != nil {
		return nil, err
	}

	err := s.repo.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	response := dto.Response{
		ID:        user.ID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil

}
