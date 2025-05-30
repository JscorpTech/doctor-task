package usecase

import (
	"errors"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/pkg/utils"
)

type RegisterUsecase struct {
	UserRepository domain.UserRepository
}

func NewRegisterUsecase(userRepository domain.UserRepository) domain.RegisterUsecase {
	return &RegisterUsecase{
		UserRepository: userRepository,
	}
}

func (r *RegisterUsecase) CreateUser(phone, firstName, lastName, password string) (*models.User, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := r.UserRepository.Create(phone, firstName, lastName, passwordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RegisterUsecase) CreateUserIfNotExist(phone, firstName, lastName, password string) (*models.User, error) {
	if r.UserRepository.IsAlready(phone) {
		return nil, errors.New("user already exists")
	}
	user, err := r.CreateUser(phone, firstName, lastName, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
