package persistence

import (
	"errors"
	"time"

	"main/src/features/users/domain"
	"main/src/features/users/domain/entities"

	"gorm.io/gorm"
)

// UserModel is the persistence representation mapped by GORM.
type UserModel struct {
	ID            uint   `gorm:"primaryKey"`
	FullName      string `gorm:"not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"`
	Phone         string
	Address       string
	ProfilePicURL string
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
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

// GormUserRepository implements UserRepository with GORM.
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
