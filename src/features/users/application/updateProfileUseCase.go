package application

import (
	"main/src/features/users/domain"
	"main/src/features/users/domain/entities"
	"time"
)

type UpdateProfileInput struct {
	FullName      *string
	Phone         *string
	Address       *string
	ProfilePicURL *string
}

type UpdateProfileUseCase struct {
	repo domain.UserRepository
}

func NewUpdateProfileUseCase(repo domain.UserRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{repo: repo}
}

func (uc *UpdateProfileUseCase) Execute(userID uint, input UpdateProfileInput) (*entities.User, error) {
	user, err := uc.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if input.FullName != nil {
		user.FullName = *input.FullName
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.Address != nil {
		user.Address = *input.Address
	}
	if input.ProfilePicURL != nil {
		user.ProfilePicURL = *input.ProfilePicURL
	}
	user.UpdatedAt = time.Now()

	if err := uc.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
