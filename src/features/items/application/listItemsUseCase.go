package application

import (
	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
	"strings"
)

type ListItemsInput struct {
	Category string
	Search   string
	Limit    int
	Offset   int
}

type ListItemsUseCase struct {
	repo domain.ItemRepository
}

func NewListItemsUseCase(repo domain.ItemRepository) *ListItemsUseCase {
	return &ListItemsUseCase{repo: repo}
}

func (uc *ListItemsUseCase) Execute(input ListItemsInput) ([]entities.Item, error) {
	if input.Category != "" {
		input.Category = strings.ToLower(strings.TrimSpace(input.Category))
	}
	if input.Limit <= 0 {
		input.Limit = 20
	}
	if input.Offset < 0 {
		input.Offset = 0
	}
	return uc.repo.List(input.Category, input.Search, input.Limit, input.Offset)
}
