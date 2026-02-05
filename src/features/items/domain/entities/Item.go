package entities

import "time"

type Item struct {
	ID           uint
	Title        string
	Description  string
	Price        float64
	PriceType    string
	Category     string
	CategorySlug string
	ImageURL     string
	OwnerID      uint
	IsAvailable  bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewItem(title, description string, price float64, priceType, category, categorySlug, imageURL string, ownerID uint) *Item {
	return &Item{
		Title:        title,
		Description:  description,
		Price:        price,
		PriceType:    priceType,
		Category:     category,
		CategorySlug: categorySlug,
		ImageURL:     imageURL,
		OwnerID:      ownerID,
		IsAvailable:  true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
