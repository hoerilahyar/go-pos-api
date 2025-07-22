package router

import (
	"gopos/internal/delivery/http/handler"
	"gopos/internal/delivery/http/middleware"
	"gopos/internal/repository"
	"gopos/internal/usecase"
	"gopos/pkg/casbin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(r *gin.Engine, db *gorm.DB) {

	enforcer := casbin.InitCasbin(db)

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		userRepo := repository.NewUserRepository(db)
		authUC := usecase.NewAuthUsecase(userRepo)
		authHandler := handler.NewAuthHandler(authUC)

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Optional: User routes jika sudah ada UserHandler
		userUC := usecase.NewUserUsecase(userRepo)
		userHandler := handler.NewUserHandler(userUC)
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware())
		users.Use(middleware.CasbinMiddleware(enforcer, db))
		{
			users.GET("", userHandler.List)
			users.GET("/:id", userHandler.Detail)
			users.POST("", userHandler.Create)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)

		}

		authorizeRepo := repository.NewAuthorizeRepository(db, enforcer)
		authorizeUC := usecase.NewAuthorizeUsecase(authorizeRepo)
		authorizeHandler := handler.NewAuthorizeHandler(authorizeUC)
		authorize := api.Group("/authorize")
		authorize.Use(middleware.AuthMiddleware())
		authorize.Use(middleware.CasbinMiddleware(enforcer, db))
		{
			authorize.GET("/policies", authorizeHandler.ListPolicies)
			authorize.GET("/policy/:id", authorizeHandler.ShowPolicy)
			authorize.POST("/policy", authorizeHandler.CreatePolicy)
			authorize.PUT("/policy", authorizeHandler.UpdatePolicy)
			authorize.DELETE("/policy/:id", authorizeHandler.DeletePolicy)

			authorize.PUT("/assign-role-to-user", authorizeHandler.AssignRoleToUser)
			authorize.DELETE("/revoke-role-from-user", authorizeHandler.RemoveRoleFromUser)
			authorize.PUT("/assign-permission-to-role", authorizeHandler.AssignPermissionToRole)
			authorize.DELETE("/revoke-permission-from-role", authorizeHandler.RemovePermissionFromRole)

			// authorize.GET("/permission", authorizeHandler.GetAllPermissions)
		}

		productRepo := repository.NewProductRepository(db)
		productUC := usecase.NewProductUsecase(productRepo)
		productHandler := handler.NewProductHandler(productUC)
		products := api.Group("/products")
		products.Use(middleware.AuthMiddleware())
		products.Use(middleware.CasbinMiddleware(enforcer, db))
		{
			products.GET("", productHandler.FindAll)
			products.GET("/:id", productHandler.FindByID)
			products.POST("", productHandler.Create)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)

		}

		categoryRepo := repository.NewCategoryRepository(db)
		categoryUC := usecase.NewCategoryUsecase(categoryRepo)
		categoryHandler := handler.NewCategoryHandler(categoryUC)
		category := api.Group("/category")
		category.Use(middleware.AuthMiddleware())
		category.Use(middleware.CasbinMiddleware(enforcer, db))
		{
			category.GET("", categoryHandler.FindAll)
			category.GET("/:id", categoryHandler.FindByID)
			category.POST("", categoryHandler.Create)
			category.PUT("/:id", categoryHandler.Update)
			category.DELETE("/:id", categoryHandler.Delete)

		}
	}

}
