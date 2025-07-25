package usecase

import (
	"gopos/internal/domain"
	"gopos/internal/repository"
	appErr "gopos/pkg/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductUsecase interface {
	FindPaginated(c *gin.Context) ([]domain.Product, int64, error)
	FindAll() ([]domain.Product, error)
	FindByID(id uint64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(req *domain.ProductUpdate) (*domain.Product, error)
	Delete(product *domain.Product) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) FindPaginated(c *gin.Context) ([]domain.Product, int64, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filters := map[string]interface{}{}

	// Query params: ?name=kopi&category_id=1&min_price=10000&max_price=50000
	if code := c.Query("code"); code != "" {
		filters["code"] = code
	}

	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := strconv.Atoi(categoryID); err == nil {
			filters["category_id"] = id
		}
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		if price, err := strconv.Atoi(minPrice); err == nil {
			filters["min_price"] = price
		}
	}

	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if price, err := strconv.Atoi(maxPrice); err == nil {
			filters["max_price"] = price
		}
	}

	products, total, err := u.productRepo.FindPaginatedWithFilter(page, limit, filters)
	if err != nil {
		return nil, 0, appErr.Get(appErr.ErrProductList, err)
	}

	return products, total, nil
}

func (u *productUsecase) FindAll() ([]domain.Product, error) {
	return u.productRepo.FindAll()
}

func (u *productUsecase) FindByID(id uint64) (*domain.Product, error) {

	product, err := u.productRepo.FindByID(id)
	if err != nil || product == nil {
		return nil, appErr.Get(appErr.ErrProductShow, err)
	}
	return product, nil
}

func (u *productUsecase) Create(product *domain.Product) error {
	return u.productRepo.Create(product)
}

func (u *productUsecase) Update(req *domain.ProductUpdate) (*domain.Product, error) {
	product, err := u.productRepo.FindByID(req.ID)
	if err != nil || product == nil {
		return nil, appErr.Get(appErr.ErrProductShow, err)
	}

	if req.Code != "" {
		product.Code = req.Code
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price != nil {
		product.Price = req.Price
	}

	return product, u.productRepo.Update(product)
}

func (u *productUsecase) Delete(product *domain.Product) error {
	return u.productRepo.Delete(product)
}
