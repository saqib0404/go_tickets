package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorUserAlreadyExists = errors.New("user with this email already exists")

type Repository interface {
	CreateUser(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorUserAlreadyExists
		}
		return result.Error
	}
	return nil
}
