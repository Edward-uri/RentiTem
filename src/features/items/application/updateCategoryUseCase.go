package application

import (
	"errors"
	"strings"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
)

// UpdateCategoryUseCase updates category name (and slug).
type UpdateCategoryUseCase struct {
	repo domain.CategoryRepository
}

func NewUpdateCategoryUseCase(repo domain.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{repo: repo}
}

func (uc *UpdateCategoryUseCase) Execute(id uint, name string) (*entities.Category, error) {
	clean := strings.TrimSpace(name)
	if clean == "" {
		return nil, errors.New("name is required")
	}
	return uc.repo.Update(id, clean)
}
