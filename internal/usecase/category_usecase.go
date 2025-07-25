package usecase

import (
	"gopos/internal/domain"
	"gopos/internal/repository"
	appErr "gopos/pkg/errors"
)

type CategoryUsecase interface {
	FindPaginated(page, limit int) ([]domain.Category, int64, error)
	FindAll() ([]domain.Category, error)
	FindByID(id uint64) (*domain.Category, error)
	Create(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(category *domain.Category) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (u *categoryUsecase) FindPaginated(page, limit int) ([]domain.Category, int64, error) {
	return u.categoryRepo.FindPaginated(page, limit)
}

func (u *categoryUsecase) FindAll() ([]domain.Category, error) {
	return u.categoryRepo.FindAll()
}

func (u *categoryUsecase) FindByID(id uint64) (*domain.Category, error) {
	category, err := u.categoryRepo.FindByID(id)
	if err != nil || category == nil {
		return nil, appErr.Get(appErr.ErrCategoryShow, err)
	}
	return category, nil
}

func (u *categoryUsecase) Create(category *domain.Category) error {
	return u.categoryRepo.Create(category)
}

func (u *categoryUsecase) Update(category *domain.Category) error {
	category, err := u.categoryRepo.FindByID(category.ID)
	if err != nil || category == nil {
		return appErr.Get(appErr.ErrCategoryShow, err)
	}
	return u.categoryRepo.Update(category)
}

func (u *categoryUsecase) Delete(category *domain.Category) error {
	category, err := u.categoryRepo.FindByID(category.ID)
	if err != nil || category == nil {
		return appErr.Get(appErr.ErrCategoryShow, err)
	}
	return u.categoryRepo.Delete(category)
}
