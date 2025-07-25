package repository

import (
	"errors"
	"gopos/internal/domain"
	appError "gopos/pkg/errors"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

type AuthorizeRepository interface {
	GetAllPolicies() ([]domain.CasbinRule, error)
	CreatePolicy(req domain.CasbinRule) (bool, error)
	ShowPolicies(id int) (domain.CasbinRule, error)
	UpdatePolicy(req domain.CasbinRule) (bool, error)
	DeletePolicy(id int) (bool, error)

	// Permission / policy management
	GetPolicies() ([][]string, error)
	AddPolicy(role, object, action string) (bool, error)
	RemovePolicy(role, object, action string) (bool, error)

	// User-role (grouping) management
	GetUserRoles(userID string) ([]string, error)
	AddUserRole(userID, role string) (bool, error)
	AssignUserRole(userID, role string) (bool, error)
	RemoveUserRole(userID, role string) (bool, error)

	// Authorization check
	CheckPermission(userID, object, action string) (bool, error)
}

type authorizeRepository struct {
	db       *gorm.DB
	enforcer *casbin.SyncedEnforcer
}

func NewAuthorizeRepository(db *gorm.DB, enforcer *casbin.SyncedEnforcer) AuthorizeRepository {
	return &authorizeRepository{db: db, enforcer: enforcer}
}

func (r *authorizeRepository) GetAllPolicies() ([]domain.CasbinRule, error) {
	var casbinRules []domain.CasbinRule
	err := r.db.Raw("SELECT * FROM casbin_rule WHERE ptype = 'p'").Scan(&casbinRules).Error
	if err != nil {
		return nil, appError.ParseMySQLError(err)
	}

	return casbinRules, nil
}

func (r *authorizeRepository) ShowPolicies(id int) (domain.CasbinRule, error) {
	var casbinRules domain.CasbinRule
	err := r.db.Raw("SELECT * FROM casbin_rule WHERE id = ?", id).Scan(&casbinRules).Error
	if err != nil {
		return casbinRules, appError.ParseMySQLError(err)
	}

	return casbinRules, nil
}

func (r *authorizeRepository) CreatePolicy(req domain.CasbinRule) (bool, error) {
	err := r.db.Exec(`
		INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		"p", req.V0, req.V1, req.V2, req.V3, req.V4, req.V5,
	).Error

	if err != nil {
		return false, appError.ParseMySQLError(err)
	}
	return true, nil
}

func (r *authorizeRepository) UpdatePolicy(req domain.CasbinRule) (bool, error) {
	err := r.db.Exec(`
			UPDATE casbin_rule
			SET ptype = ?, v0 = ?, v1 = ?, v2 = ?, v3 = ?, v4 = ?, v5 = ?
			WHERE id = ?
		`, "p", req.V0, req.V1, req.V2, req.V3, req.V4, req.V5, req.ID).Error

	if err != nil {
		return false, appError.ParseMySQLError(err)
	}
	return true, nil
}

// Delete policy by ID
func (r *authorizeRepository) DeletePolicy(id int) (bool, error) {
	result := r.db.Exec("DELETE FROM casbin_rule WHERE id = ?", id)
	if result.Error != nil {
		return false, appError.ParseMySQLError(result.Error)
	}
	if result.RowsAffected == 0 {
		return false, errors.New("no record found")
	}
	return true, nil
}

func (r *authorizeRepository) GetPolicies() ([][]string, error) {
	return r.enforcer.GetPolicy()
}

func (r *authorizeRepository) AddPolicy(role, object, action string) (bool, error) {
	return r.enforcer.AddPolicy(role, object, action)
}

func (r *authorizeRepository) RemovePolicy(role, object, action string) (bool, error) {
	return r.enforcer.RemovePolicy(role, object, action)
}

func (r *authorizeRepository) GetUserRoles(userID string) ([]string, error) {
	return r.enforcer.GetRolesForUser(userID)
}

func (r *authorizeRepository) AddUserRole(userID, role string) (bool, error) {
	return r.enforcer.AddGroupingPolicy(userID, role)
}

func (r *authorizeRepository) AssignUserRole(userID, role string) (bool, error) {
	return r.enforcer.AddGroupingPolicy(userID, role)
}

func (r *authorizeRepository) RemoveUserRole(userID, role string) (bool, error) {
	return r.enforcer.RemoveGroupingPolicy(userID, role)
}

func (r *authorizeRepository) CheckPermission(userID, object, action string) (bool, error) {
	return r.enforcer.Enforce(userID, object, action)
}
