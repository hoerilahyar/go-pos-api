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

type ProductHandler struct {
	productUC usecase.ProductUsecase
}

func NewProductHandler(productUC usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUC: productUC}
}

func (h *ProductHandler) FindAll(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, total, err := h.productUC.FindPaginated(c)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "List Product successful", products, page, limit, total)
}

func (h *ProductHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	product, err := h.productUC.FindByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "List Detail successful", product)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &product); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}
	if err := h.productUC.Create(&product); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Create Product successful", product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	var product domain.ProductUpdate

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, errors.New("Invalid JSON body"))
		return
	}

	product.ID = id

	updateData, err := h.productUC.Update(&product)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Update Product successful", updateData)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c, errors.New("Invalid ID"))
		return
	}

	product, err := h.productUC.FindByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}
	if product == nil {
		response.Error(c, errors.New("Product not found"))
		return
	}

	if err := h.productUC.Delete(product); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Delete Product successful")
}
