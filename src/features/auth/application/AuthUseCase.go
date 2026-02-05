package application

import (
	"errors"
	"strings"

	"main/src/features/auth/domain"
	"main/src/features/users/domain/entities"
)

type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hashed, plain string) bool
}

type TokenService interface {
	Generate(userID uint, email, role string) (string, error)
}

type RegisterInput struct {
	FullName string
	Email    string
	Password string
	Phone    string
	Address  string
	Role     string
}

type LoginInput struct {
	Email    string
	Password string
}

type AuthUseCase struct {
	repo        domain.AuthRepository
	pwdService  PasswordService
	tokenSource TokenService
}

func NewAuthUseCase(repo domain.AuthRepository, pwd PasswordService, token TokenService) *AuthUseCase {
	return &AuthUseCase{repo: repo, pwdService: pwd, tokenSource: token}
}

func (uc *AuthUseCase) Register(input RegisterInput) (string, *entities.User, error) {
	input.Email = strings.TrimSpace(input.Email)
	input.FullName = strings.TrimSpace(input.FullName)
	input.Role = strings.ToLower(strings.TrimSpace(input.Role))

	if input.Email == "" || input.Password == "" || input.FullName == "" || input.Phone == "" || input.Address == "" {
		return "", nil, errors.New("missing required fields")
	}
	if input.Role == "" {
		input.Role = "user"
	}
	if input.Role != "user" && input.Role != "superadmin" {
		return "", nil, errors.New("invalid role")
	}
	if len(input.Password) < 6 {
		return "", nil, errors.New("password too short")
	}

	existing, err := uc.repo.FindByEmail(input.Email)
	if err != nil {
		return "", nil, err
	}
	if existing != nil {
		return "", nil, errors.New("user already exists")
	}

	hash, err := uc.pwdService.Hash(input.Password)
	if err != nil {
		return "", nil, err
	}

	user := entities.NewUser(input.FullName, input.Email, hash, input.Phone, input.Address, "", input.Role)
	if err := uc.repo.Create(user); err != nil {
		return "", nil, err
	}

	token, err := uc.tokenSource.Generate(user.ID, user.Email, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (uc *AuthUseCase) Login(input LoginInput) (string, *entities.User, error) {
	input.Email = strings.TrimSpace(input.Email)
	if input.Email == "" || input.Password == "" {
		return "", nil, errors.New("missing credentials")
	}

	user, err := uc.repo.FindByEmail(input.Email)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !uc.pwdService.Compare(user.Password, input.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := uc.tokenSource.Generate(user.ID, user.Email, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
