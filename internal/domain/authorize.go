package domain

type CasbinRule struct {
	ID    *uint  `gorm:"primaryKey" json:"id"`
	PType string `gorm:"column:ptype" json:"p_type" `
	V0    string `gorm:"column:v0" json:"v0" binding:"required"` // Subject (role or user ID)
	V1    string `gorm:"column:v1" json:"v1"`                    // Object (route/resource)
	V2    string `gorm:"column:v2" json:"v2"`                    // Action (GET, POST, etc)
	V3    string `gorm:"column:v3" json:"v3"`
	V4    string `gorm:"column:v4" json:"v4"`
	V5    string `gorm:"column:v5" json:"v5"`
}

type AssignPermissionRequest struct {
	Role   string `json:"role" binding:"required"`
	Object string `json:"object" binding:"required"` // route path, e.g. "/users"
	Action string `json:"action" binding:"required"` // e.g. GET, POST, etc
}

// Request untuk assign role ke user
type AssignRoleRequest struct {
	Role   string `json:"role" binding:"required"` // super-admin
	Object string `json:"object" default:"*"`      // * (optional)
	Action string `json:"action" default:"*"`      // * (optional)
}

type PolicyRequest struct {
	Role   string `json:"role" binding:"required"` // super-admin
	Object string `json:"object" default:"*"`      // * (optional)
	Action string `json:"action" default:"*"`      // * (optional)
}

// Request untuk assign user ke role (sama dengan AssignRoleRequest)
type AssignUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}

type PolicyResponse struct {
	Policies [][]string `json:"policies"`
}

type RoleResponse struct {
	Roles []string `json:"roles"`
}

type UserRoleResponse struct {
	Roles []string `json:"roles"`
}

type AssignUserRoleRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}
