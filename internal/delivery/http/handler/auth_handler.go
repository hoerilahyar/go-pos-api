package handler

import (
	"gopos/internal/domain"
	"gopos/internal/usecase"
	"gopos/pkg/errors"
	"gopos/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.ErrValidation)
		return
	}

	user, err := h.authUC.Register(&req)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Register successful", user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err)
		return
	}

	login, err := h.authUC.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Login successful", login)
}

func (h *AuthHandler) AuthInfo(c *gin.Context) {

	// response.Success(c, "Login successful", info)
}
