package application

import (
	"errors"

	"main/src/features/users/domain"
	"main/src/features/users/domain/entities"
)

type GetProfileUseCase struct {
	repo domain.UserRepository
}

func NewGetProfileUseCase(repo domain.UserRepository) *GetProfileUseCase {
	return &GetProfileUseCase{repo: repo}
}

func (uc *GetProfileUseCase) Execute(userID uint) (*entities.User, error) {
	user, err := uc.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
