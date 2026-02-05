package application

import (
	"errors"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
	usersDomain "main/src/features/users/domain"
)

type ItemWithOwner struct {
	Item       *entities.Item
	OwnerName  string
	OwnerPhone string
}

type GetItemUseCase struct {
	repo     domain.ItemRepository
	userRepo usersDomain.UserRepository
}

func NewGetItemUseCase(repo domain.ItemRepository, userRepo usersDomain.UserRepository) *GetItemUseCase {
	return &GetItemUseCase{repo: repo, userRepo: userRepo}
}

func (uc *GetItemUseCase) Execute(id uint) (*ItemWithOwner, error) {
	item, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("item not found")
	}

	owner, err := uc.userRepo.FindByID(item.OwnerID)
	if err != nil {
		return nil, err
	}

	ownerName := ""
	ownerPhone := ""
	if owner != nil {
		ownerName = owner.FullName
		ownerPhone = owner.Phone
	}

	return &ItemWithOwner{Item: item, OwnerName: ownerName, OwnerPhone: ownerPhone}, nil
}
