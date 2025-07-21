package handler

import (
	"errors"
	"gopos/internal/domain"
	"gopos/internal/usecase"
	"gopos/pkg/response"
	"gopos/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthorizeHandler struct {
	authUC usecase.AuthorizeUsecase
}

func NewAuthorizeHandler(authUC usecase.AuthorizeUsecase) *AuthorizeHandler {
	return &AuthorizeHandler{authUC: authUC}
}

func (h *AuthorizeHandler) ListPolicies(c *gin.Context) {
	policies, err := h.authUC.GetPolicies()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "List Policies successful", policies)
}

func (h *AuthorizeHandler) ShowPolicy(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := utils.StrToInt(idParam)
	rule, err := h.authUC.ShowPolicy(id)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Show Policy successful", rule)
}

func (h *AuthorizeHandler) CreatePolicy(c *gin.Context) {
	var req domain.CasbinRule

	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	user, err := h.authUC.CreatePolicy(req)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Create Policy successful", user)
}

func (h *AuthorizeHandler) UpdatePolicy(c *gin.Context) {
	var req domain.CasbinRule

	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	if req.ID == nil {
		response.Error(c, errors.New("Invalid Payload"), []string{"ID required"})
		return
	}

	user, err := h.authUC.UpdatePolicy(req)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Update Policy successful", user)
}

func (h *AuthorizeHandler) DeletePolicy(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := utils.StrToInt(idParam)
	rule, err := h.authUC.DeletePolicy(id)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Delete Policy successful", rule)
}

func (h *AuthorizeHandler) AssignRoleToUser(c *gin.Context) {
	var req domain.AssignUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	ok, err := h.authUC.AssignUserRole(req.UserID, req.Role)
	if err != nil || !ok {
		response.Error(c, err)
		return
	}

	response.Success(c, "Role assigned")
}

func (h *AuthorizeHandler) RemoveRoleFromUser(c *gin.Context) {
	var req domain.AssignUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	ok, err := h.authUC.RemoveUserRole(req.UserID, req.Role)

	if err != nil || !ok {
		response.Error(c, err)
		return
	}

	response.Success(c, "Role removed")
}

func (h *AuthorizeHandler) AssignPermissionToRole(c *gin.Context) {
	var req domain.AssignPermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	_, err := h.authUC.AssignPermission(&req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Permission assigned")
}

func (h *AuthorizeHandler) RemovePermissionFromRole(c *gin.Context) {
	var req domain.AssignPermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		var errData []string
		if validation := utils.HandleValidationError(err, &req); len(validation) > 0 {
			errData = validation
			err = errors.New("Invalid Payload")
		}
		response.Error(c, err, errData)
		return
	}

	_, err := h.authUC.RemovePermission(&req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Permission removed")
}

// func (h *AuthorizeHandler) GetAllPermissions(c *gin.Context) {
// 	policies, err := h.authUC.GetPolicies()
// 	if err != nil {
// 		response.Error(c, err)
// 		return
// 	}

// 	response.Success(c, "List of Policies", policies)
// }
