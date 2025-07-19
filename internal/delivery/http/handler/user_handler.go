package handler

import (
	"gopos/internal/domain"
	"gopos/internal/usecase"
	"gopos/pkg/response"
	"gopos/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) List(c *gin.Context) {

	users, err := h.userUC.List()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "List User successful", users)
}

func (h *UserHandler) Detail(c *gin.Context) {
	userId := c.Param("id")

	uintUserId, err := utils.StrToUint(userId)
	if err != nil {
		response.Error(c, err)
		return
	}

	user, err := h.userUC.Detail(uintUserId)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Get User successful", user)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req domain.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err)
		return
	}

	user, err := h.userUC.Create(&req)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Create User successful", user)
}

func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := utils.StrToUint(idParam)
	if err != nil {
		response.Error(c, err)
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, err)
		return
	}

	user.ID = idUint // set ID dari URL ke struct user

	updatedUser, err := h.userUC.Update(&user)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "User updated successfully", updatedUser)
}

func (h *UserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, err)
		return
	}

	err = h.userUC.Delete(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User soft deleted"})
}
