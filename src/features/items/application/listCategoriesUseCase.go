package application

import (
	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
)

type ListCategoriesUseCase struct {
	repo domain.CategoryRepository
}

func NewListCategoriesUseCase(repo domain.CategoryRepository) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{repo: repo}
}

func (uc *ListCategoriesUseCase) Execute() ([]entities.Category, error) {
	return uc.repo.List()
}
