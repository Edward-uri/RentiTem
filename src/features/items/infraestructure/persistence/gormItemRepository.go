package persistence

import (
	"errors"
	"strings"
	"time"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"

	"gorm.io/gorm"
)

type ItemModel struct {
	ID           uint    `gorm:"primaryKey"`
	Title        string  `gorm:"not null"`
	Description  string  `gorm:"type:text;not null"`
	Price        float64 `gorm:"not null"`
	PriceType    string  `gorm:"not null"`
	Category     string  `gorm:"not null"`
	CategorySlug string  `gorm:"not null;default:'';index"`
	ImageURL     string  `gorm:"not null"`
	OwnerID      uint    `gorm:"not null"`
	IsAvailable  bool    `gorm:"not null;default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func slugifyCategory(name string) string {
	slug := strings.TrimSpace(strings.ToLower(name))
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func toModel(e *entities.Item) ItemModel {
	slug := e.CategorySlug
	if slug == "" {
		slug = slugifyCategory(e.Category)
	}
	return ItemModel{
		ID:           e.ID,
		Title:        e.Title,
		Description:  e.Description,
		Price:        e.Price,
		PriceType:    e.PriceType,
		Category:     e.Category,
		CategorySlug: slug,
		ImageURL:     e.ImageURL,
		OwnerID:      e.OwnerID,
		IsAvailable:  e.IsAvailable,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func toEntity(m ItemModel) entities.Item {
	return entities.Item{
		ID:           m.ID,
		Title:        m.Title,
		Description:  m.Description,
		Price:        m.Price,
		PriceType:    m.PriceType,
		Category:     m.Category,
		CategorySlug: m.CategorySlug,
		ImageURL:     m.ImageURL,
		OwnerID:      m.OwnerID,
		IsAvailable:  m.IsAvailable,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

type GormItemRepository struct {
	db *gorm.DB
}

var _ domain.ItemRepository = (*GormItemRepository)(nil)

func NewGormItemRepository(db *gorm.DB) *GormItemRepository {
	return &GormItemRepository{db: db}
}

func (r *GormItemRepository) Create(item *entities.Item) error {
	m := toModel(item)
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	item.ID = m.ID
	return nil
}

func (r *GormItemRepository) FindByID(id uint) (*entities.Item, error) {
	var m ItemModel
	if err := r.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	entity := toEntity(m)
	return &entity, nil
}

func (r *GormItemRepository) List(category, search string, limit, offset int) ([]entities.Item, error) {
	var models []ItemModel
	q := r.db.Model(&ItemModel{}).Where("is_available = ?", true)
	if category != "" {
		slug := strings.ToLower(category)
		q = q.Where("category_slug = ? OR (category_slug = '' AND LOWER(category) = ?)", slug, slug)
	}
	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		q = q.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", like, like)
	}
	if err := q.Limit(limit).Offset(offset).Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}
	items := make([]entities.Item, 0, len(models))
	for _, m := range models {
		items = append(items, toEntity(m))
	}
	return items, nil
}

func (r *GormItemRepository) Update(item *entities.Item) error {
	m := toModel(item)
	return r.db.Model(&ItemModel{}).Where("id = ?", item.ID).Updates(m).Error
}

func (r *GormItemRepository) Delete(id uint) error {
	return r.db.Delete(&ItemModel{}, id).Error
}
