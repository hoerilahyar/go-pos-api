package repository

import (
	"errors"
	"gopos/internal/domain"
	appError "gopos/pkg/errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindPaginated(page, limit int) ([]domain.Category, int64, error)
	FindAll() ([]domain.Category, error)
	FindByID(id uint64) (*domain.Category, error)
	Create(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(category *domain.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) FindPaginated(page, limit int) ([]domain.Category, int64, error) {
	var categories []domain.Category
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Model(&domain.Category{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	if err := r.db.Where("deleted_at IS NULL").Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	return categories, total, nil
}

func (r *categoryRepository) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Where("deleted_at IS NULL").Find(&categories).Error
	return categories, appError.ParseMySQLError(err)
}

func (r *categoryRepository) FindByID(id uint64) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Where("deleted_at IS NULL").First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, appError.ParseMySQLError(err)
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *domain.Category) error {
	err := r.db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Update(category *domain.Category) error {
	if err := r.db.Model(&domain.User{}).Where("id = ? AND deleted_at IS NULL", category.ID).Updates(category).Error; err != nil {
		return appError.ParseMySQLError(err)
	}
	return nil
}

func (r *categoryRepository) Delete(category *domain.Category) error {
	err := r.db.Delete(category).Error
	if err != nil {
		return appError.ParseMySQLError(err)
	}
	return nil
}
