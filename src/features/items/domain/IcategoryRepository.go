package domain

import "main/src/features/items/domain/entities"

type CategoryRepository interface {
	Resolve(input string) (*entities.Category, error)
	EnsureDefaults(defaults []entities.Category) error
	List() ([]entities.Category, error)
	Create(name string) (*entities.Category, error)
	Update(id uint, name string) (*entities.Category, error)
	Delete(id uint) error
}
