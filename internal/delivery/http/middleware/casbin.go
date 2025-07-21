package middleware

import (
	"errors"
	"fmt"
	"gopos/pkg/response"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CasbinMiddleware(e *casbin.SyncedEnforcer, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		obj := c.FullPath()
		act := c.Request.Method

		userID, exists := c.Get("user_id")
		if !exists {
			response.Error(c, errors.New("unauthorized"))
			c.Abort()
			return
		}

		sub := fmt.Sprintf("%v", userID)

		// fmt.Println("[DEBUG] Enforcing sub =", sub, "obj =", obj, "act =", act)
		// policies, _ := e.GetPolicy()
		// fmt.Println("[DEBUG] Policies:", policies)

		// groupingPolicies, _ := e.GetGroupingPolicy()
		// fmt.Println("[DEBUG] Grouping Policies:", groupingPolicies)

		allowed, err := e.Enforce(sub, obj, act)
		if err != nil {
			response.Error(c, errors.New("authorization error"))
			c.Abort()
			return
		}

		if !allowed {
			response.Error(c, errors.New("forbidden"))
			c.Abort()
			return
		}

		c.Next()
	}
}
