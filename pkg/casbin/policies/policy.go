package policies

import (
	"github.com/casbin/casbin/v2"
)

// SetupPolicies sets and saves all the required policies.
func SetupPolicies(enforcer *casbin.SyncedEnforcer) error {
	enforcer.AddPolicy("super-admin", "*", "*")
	return enforcer.SavePolicy()
}
