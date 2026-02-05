package persistence

import (
	"strings"

	"main/src/features/items/domain"
	"main/src/features/items/domain/entities"

	"gorm.io/gorm"
)

type CategoryModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Slug string `gorm:"unique;not null;index"`
}

func toCategoryEntity(m CategoryModel) entities.Category {
	return entities.Category{ID: m.ID, Name: m.Name, Slug: m.Slug}
}

type GormCategoryRepository struct {
	db *gorm.DB
}

var _ domain.CategoryRepository = (*GormCategoryRepository)(nil)

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{db: db}
}

func slugify(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func (r *GormCategoryRepository) EnsureDefaults(defaults []entities.Category) error {
	for _, c := range defaults {
		slug := slugify(c.Name)
		model := CategoryModel{Name: c.Name, Slug: slug}
		if err := r.db.Where(CategoryModel{Slug: slug}).FirstOrCreate(&model).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *GormCategoryRepository) Create(name string) (*entities.Category, error) {
	slug := slugify(name)
	model := CategoryModel{Name: name, Slug: slug}
	if err := r.db.Create(&model).Error; err != nil {
		return nil, err
	}
	entity := toCategoryEntity(model)
	return &entity, nil
}

func (r *GormCategoryRepository) Update(id uint, name string) (*entities.Category, error) {
	slug := slugify(name)
	var model CategoryModel
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	model.Name = name
	model.Slug = slug
	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}
	entity := toCategoryEntity(model)
	return &entity, nil
}

func (r *GormCategoryRepository) Delete(id uint) error {
	result := r.db.Delete(&CategoryModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormCategoryRepository) Resolve(input string) (*entities.Category, error) {
	clean := strings.TrimSpace(input)
	if clean == "" {
		return r.getDefault()
	}
	cleanSlug := slugify(clean)

	var model CategoryModel
	if err := r.db.Where("slug = ? OR LOWER(name) = ?", cleanSlug, cleanSlug).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.getDefault()
		}
		return nil, err
	}
	entity := toCategoryEntity(model)
	return &entity, nil
}

func (r *GormCategoryRepository) getDefault() (*entities.Category, error) {
	var model CategoryModel
	if err := r.db.Where("slug = ?", "otro").First(&model).Error; err != nil {
		return nil, err
	}
	entity := toCategoryEntity(model)
	return &entity, nil
}

func (r *GormCategoryRepository) List() ([]entities.Category, error) {
	var models []CategoryModel
	if err := r.db.Order("name ASC").Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]entities.Category, 0, len(models))
	for _, m := range models {
		res = append(res, toCategoryEntity(m))
	}
	return res, nil
}
