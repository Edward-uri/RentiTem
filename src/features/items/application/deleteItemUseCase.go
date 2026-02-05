package application

import (
	"errors"

	"main/src/features/items/domain"
)

type DeleteItemUseCase struct {
	repo domain.ItemRepository
}

func NewDeleteItemUseCase(repo domain.ItemRepository) *DeleteItemUseCase {
	return &DeleteItemUseCase{repo: repo}
}

func (uc *DeleteItemUseCase) Execute(id uint, ownerID uint) error {
	item, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("item not found")
	}
	if item.OwnerID != ownerID {
		return errors.New("forbidden")
	}
	return uc.repo.Delete(id)
}
