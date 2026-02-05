package application

import "main/src/features/items/domain"

type DeleteCategoryUseCase struct {
	repo domain.CategoryRepository
}

func NewDeleteCategoryUseCase(repo domain.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{repo: repo}
}

func (uc *DeleteCategoryUseCase) Execute(id uint) error {
	return uc.repo.Delete(id)
}
