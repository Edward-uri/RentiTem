package application

import (
	"errors"
	"strings"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
)

// CreateCategoryUseCase handles creation of a new category.
type CreateCategoryUseCase struct {
	repo domain.CategoryRepository
}

func NewCreateCategoryUseCase(repo domain.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{repo: repo}
}

func (uc *CreateCategoryUseCase) Execute(name string) (*entities.Category, error) {
	clean := strings.TrimSpace(name)
	if clean == "" {
		return nil, errors.New("name is required")
	}
	return uc.repo.Create(clean)
}
