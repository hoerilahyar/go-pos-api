package repository

import (
	"errors"

	"gopos/internal/domain"
	appError "gopos/pkg/errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) List() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, appError.ParseMySQLError(err)
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found, return nil
		}
		return nil, appError.ParseMySQLError(err)
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found, return nil
		}
		return nil, appError.ParseMySQLError(err)
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
		return nil, appError.ParseMySQLError(err)
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
		return nil, appError.ParseMySQLError(err)
	}
	return &user, nil
}

func (r *userRepository) Save(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, appError.ParseMySQLError(err)
	}
	return user, nil
}

// ✅ Update user data (name, email, role, etc.)
func (r *userRepository) Update(user *domain.User) (*domain.User, error) {
	updateData := map[string]interface{}{}

	if user.Username != "" {
		updateData["username"] = user.Username
	}

	if user.Name != "" {
		updateData["name"] = user.Name
	}

	if user.Email != "" {
		updateData["email"] = user.Email
	}

	if user.Password != "" {
		updateData["password"] = user.Password
	}

	if err := r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(updateData).Error; err != nil {
		return nil, appError.ParseMySQLError(err)
	}

	// Ambil kembali user setelah update
	var updatedUser domain.User
	if err := r.db.First(&updatedUser, user.ID).Error; err != nil {
		return nil, appError.ParseMySQLError(err)
	}

	return &updatedUser, nil
}

// ✅ Soft delete user
func (r *userRepository) Delete(user *domain.User) error {
	err := r.db.Delete(user).Error
	if err != nil {
		return appError.ParseMySQLError(err)
	}
	return nil
}
