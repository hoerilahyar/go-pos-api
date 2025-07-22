package handler

import (
	"errors"
	"gopos/internal/domain"
	"gopos/internal/usecase"
	"gopos/pkg/response"
	"gopos/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUC usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUC usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUC: categoryUC}
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	categories, total, err := h.categoryUC.FindPaginated(page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "List Category successful", categories, page, limit, total)
}

func (h *CategoryHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	category, err := h.categoryUC.FindByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Detail Category successful", category)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &category); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}
	if err := h.categoryUC.Create(&category); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Create Category successful", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &category); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}
	category.ID = id

	if err := h.categoryUC.Update(&category); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Update Category successful", category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &category); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}
	category.ID = id

	if err := h.categoryUC.Delete(&category); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Delete Category successful")
}
