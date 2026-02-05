package domain

import "main/src/features/items/domain/entities"

type ItemRepository interface {
	Create(item *entities.Item) error
	FindByID(id uint) (*entities.Item, error)
	List(category, search string, limit, offset int) ([]entities.Item, error)
	Update(item *entities.Item) error
	Delete(id uint) error
}
