package application

import (
	"errors"
	"strings"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"
)

type StorageService interface {
	Save(fileBytes []byte, filename string) (string, error)
}

type CreateItemInput struct {
	Title       string
	Description string
	Price       float64
	PriceType   string
	Category    string
	ImageURL    string
	OwnerID     uint
}

type CreateItemUseCase struct {
	repo    domain.ItemRepository
	catRepo domain.CategoryRepository
}

func NewCreateItemUseCase(repo domain.ItemRepository, catRepo domain.CategoryRepository) *CreateItemUseCase {
	return &CreateItemUseCase{repo: repo, catRepo: catRepo}
}

func (uc *CreateItemUseCase) Execute(input CreateItemInput) (*entities.Item, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	input.Category = strings.TrimSpace(input.Category)
	input.PriceType = strings.TrimSpace(strings.ToUpper(input.PriceType))

	if input.Title == "" || input.Description == "" || input.Category == "" || input.Price == 0 || input.PriceType == "" || input.ImageURL == "" || input.OwnerID == 0 {
		return nil, errors.New("missing required fields")
	}
	if input.PriceType != "POR_HORA" && input.PriceType != "POR_DIA" {
		return nil, errors.New("invalid price_type")
	}

	category, err := uc.catRepo.Resolve(input.Category)
	if err != nil {
		return nil, err
	}

	item := entities.NewItem(input.Title, input.Description, input.Price, input.PriceType, category.Name, category.Slug, input.ImageURL, input.OwnerID)
	if err := uc.repo.Create(item); err != nil {
		return nil, err
	}
	return item, nil
}
