package persistence

import (
	"errors"
	"time"

	"main/src/features/users/domain"
	"main/src/features/users/domain/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	ID            uint   `gorm:"primaryKey"`
	FullName      string `gorm:"not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"`
	Phone         string
	Address       string
	ProfilePicURL string
	Role          string `gorm:"not null;default:user"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func toModel(e *entities.User) UserModel {
	return UserModel{
		ID:            e.ID,
		FullName:      e.FullName,
		Email:         e.Email,
		Password:      e.Password,
		Phone:         e.Phone,
		Address:       e.Address,
		ProfilePicURL: e.ProfilePicURL,
		Role:          e.Role,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func toEntity(m UserModel) *entities.User {
	return &entities.User{
		ID:            m.ID,
		FullName:      m.FullName,
		Email:         m.Email,
		Password:      m.Password,
		Phone:         m.Phone,
		Address:       m.Address,
		ProfilePicURL: m.ProfilePicURL,
		Role:          m.Role,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

type GormUserRepository struct {
	db *gorm.DB
}

var _ domain.UserRepository = (*GormUserRepository)(nil)

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *entities.User) error {
	model := toModel(user)
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}
	user.ID = model.ID
	return nil
}

func (r *GormUserRepository) FindByEmail(email string) (*entities.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toEntity(model), nil
}

func (r *GormUserRepository) FindByID(id uint) (*entities.User, error) {
	var model UserModel
	if err := r.db.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toEntity(model), nil
}

func (r *GormUserRepository) Update(user *entities.User) error {
	model := toModel(user)
	return r.db.Model(&UserModel{}).Where("id = ?", user.ID).Updates(model).Error
}
