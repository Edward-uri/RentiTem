package application

import (
	"errors"

	"main/src/features/users/domain"
	"main/src/features/users/domain/entities"
)

type CreateUserInput struct {
	FullName string
	Email    string
	Password string
	Phone    string
	Address  string
	Profile  string
	Role     string
}

type CreateUserUseCase struct {
	repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInput) (*entities.User, error) {
	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email and password are required")
	}

	existing, err := uc.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	user := entities.NewUser(
		input.FullName,
		input.Email,
		input.Password,
		input.Phone,
		input.Address,
		input.Profile,
		input.Role,
	)

	if err := uc.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
