package usecase

import (
	"gopos/internal/domain"
	"gopos/internal/repository"
)

type AuthorizeUsecase interface {
	GetListPolicy() ([]domain.CasbinRule, error)
	ShowPolicy(id int) (domain.CasbinRule, error)
	CreatePolicy(req domain.CasbinRule) (bool, error)
	UpdatePolicy(req domain.CasbinRule) (bool, error)
	DeletePolicy(id int) (bool, error)

	GetPolicies() ([]domain.CasbinRule, error)
	AddPolicy(req *domain.PolicyRequest) (bool, error)
	RemovePolicy(role, object, action string) (bool, error)
	AssignPermission(req *domain.AssignPermissionRequest) (bool, error)
	RemovePermission(req *domain.AssignPermissionRequest) (bool, error)

	// User-Role assignment
	GetUserRoles(userID string) ([]string, error)
	AssignUserRole(userID, role string) (bool, error)
	RemoveUserRole(userID, role string) (bool, error)

	// Authorization check
	CheckPermission(userID, object, action string) (bool, error)
}

type authorizeUsecase struct {
	authorizeRepo repository.AuthorizeRepository
}

func NewAuthorizeUsecase(authorizeRepo repository.AuthorizeRepository) AuthorizeUsecase {
	return &authorizeUsecase{
		authorizeRepo: authorizeRepo,
	}
}

func (uc *authorizeUsecase) GetListPolicy() ([]domain.CasbinRule, error) {
	rule, err := uc.authorizeRepo.GetAllPolicies()

	if err != nil {
		return nil, err
	}

	return rule, nil
}

func (uc *authorizeUsecase) ShowPolicy(id int) (domain.CasbinRule, error) {

	rule, err := uc.authorizeRepo.ShowPolicies(id)

	if err != nil {
		return domain.CasbinRule{}, err
	}

	return rule, nil
}

func (uc *authorizeUsecase) GetPolicies() ([]domain.CasbinRule, error) {

	policies, err := uc.authorizeRepo.GetAllPolicies()
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (uc *authorizeUsecase) AddPolicy(req *domain.PolicyRequest) (bool, error) {
	return uc.authorizeRepo.AddPolicy(req.Role, req.Object, req.Action)
}

func (uc *authorizeUsecase) CreatePolicy(req domain.CasbinRule) (bool, error) {
	ok, err := uc.authorizeRepo.CreatePolicy(req)

	if err != nil || !ok {
		return false, err
	}

	return true, nil
}

func (uc *authorizeUsecase) UpdatePolicy(req domain.CasbinRule) (bool, error) {

	ok, err := uc.authorizeRepo.UpdatePolicy(req)

	if err != nil || !ok {
		return false, err
	}

	return true, nil
}

func (uc *authorizeUsecase) DeletePolicy(id int) (bool, error) {
	ok, err := uc.authorizeRepo.DeletePolicy(id)

	if err != nil || !ok {
		return false, err
	}

	return true, nil
}

func (uc *authorizeUsecase) RemovePolicy(role, object, action string) (bool, error) {
	return uc.authorizeRepo.RemovePolicy(role, object, action)
}

func (uc *authorizeUsecase) AssignPermission(req *domain.AssignPermissionRequest) (bool, error) {
	return uc.authorizeRepo.AddPolicy(req.Role, req.Object, req.Action)
}
func (uc *authorizeUsecase) RemovePermission(req *domain.AssignPermissionRequest) (bool, error) {
	return uc.authorizeRepo.RemovePolicy(req.Role, req.Object, req.Action)
}

func (uc *authorizeUsecase) GetUserRoles(userID string) ([]string, error) {
	return uc.authorizeRepo.GetUserRoles(userID)
}

func (uc *authorizeUsecase) AssignUserRole(userID, role string) (bool, error) {
	return uc.authorizeRepo.AssignUserRole(userID, role)
}

func (uc *authorizeUsecase) RemoveUserRole(userID, role string) (bool, error) {
	return uc.authorizeRepo.RemoveUserRole(userID, role)
}

func (uc *authorizeUsecase) CheckPermission(userID, object, action string) (bool, error) {
	return uc.authorizeRepo.CheckPermission(userID, object, action)
}
