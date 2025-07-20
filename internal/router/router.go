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
	}

}
