package repository

import (
	"errors"

	"gopos/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found, return nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmailOrUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ? OR username = ?", username, username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found, return nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

// ✅ Update user data (name, email, role, etc.)
func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

// ✅ Soft delete user
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
