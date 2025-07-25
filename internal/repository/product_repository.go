package repository

import (
	"errors"
	"gopos/internal/domain"
	appError "gopos/pkg/errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindPaginatedWithFilter(page, limit int, filters map[string]interface{}) ([]domain.Product, int64, error)
	FindPaginated(page, limit int) ([]domain.Product, int64, error)
	FindAll() ([]domain.Product, error)
	FindByID(id uint64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(category *domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindPaginated(page, limit int) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	offset := (page - 1) * limit

	if err := r.db.
		Model(&domain.Product{}).
		Where("deleted_at IS NULL").
		Count(&total).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	if err := r.db.
		Where("deleted_at IS NULL").
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	return products, total, nil
}

func (r *productRepository) FindPaginatedWithFilter(page, limit int, filters map[string]interface{}) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	offset := (page - 1) * limit

	query := r.db.Model(&domain.Product{}).Where("deleted_at IS NULL")

	for key, value := range filters {
		switch key {
		case "code":
			query = query.Where("code LIKE ?", "%"+value.(string)+"%")
		case "name":
			query = query.Where("name LIKE ?", "%"+value.(string)+"%")
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "min_price":
			query = query.Where("price >= ?", value)
		case "max_price":
			query = query.Where("price <= ?", value)
		}
	}

	// Hitung total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	// Ambil data
	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, appError.ParseMySQLError(err)
	}

	return products, total, nil
}

func (r *productRepository) FindAll() ([]domain.Product, error) {
	var product []domain.Product
	err := r.db.Where("deleted_at IS NULL").Find(&product).Error
	return product, appError.ParseMySQLError(err)
}

func (r *productRepository) FindByID(id uint64) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Where("deleted_at IS NULL").First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, appError.ParseMySQLError(err)
	}
	return &product, nil
}

func (r *productRepository) Create(product *domain.Product) error {
	err := r.db.Create(product).Error
	if err != nil {
		return appError.ParseMySQLError(err)
	}
	return nil
}

func (r *productRepository) Update(product *domain.Product) error {
	if err := r.db.Model(&domain.Product{}).Where("id = ? AND deleted_at IS NULL", product.ID).Updates(product).Error; err != nil {
		return appError.ParseMySQLError(err)
	}

	return nil
}

func (r *productRepository) Delete(product *domain.Product) error {
	err := r.db.Delete(product).Error
	if err != nil {
		return appError.ParseMySQLError(err)
	}
	return nil
}
