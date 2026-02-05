package application

import (
	"errors"
	"time"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
)

type UpdateItemInput struct {
	Title       *string
	Price       *float64
	IsAvailable *bool
}

type UpdateItemUseCase struct {
	repo domain.ItemRepository
}

func NewUpdateItemUseCase(repo domain.ItemRepository) *UpdateItemUseCase {
	return &UpdateItemUseCase{repo: repo}
}

func (uc *UpdateItemUseCase) Execute(id uint, ownerID uint, input UpdateItemInput) (*entities.Item, error) {
	item, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("item not found")
	}
	if item.OwnerID != ownerID {
		return nil, errors.New("forbidden")
	}

	if input.Title != nil {
		item.Title = *input.Title
	}
	if input.Price != nil {
		item.Price = *input.Price
	}
	if input.IsAvailable != nil {
		item.IsAvailable = *input.IsAvailable
	}
	item.UpdatedAt = time.Now()

	if err := uc.repo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}
