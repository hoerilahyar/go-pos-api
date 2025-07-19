package router

import (
	"gopos/internal/delivery/http/handler"
	"gopos/internal/repository"
	"gopos/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Optional: User routes jika sudah ada UserHandler
		// userUC := usecase.NewUserUsecase(userRepo)
		// userHandler := handler.NewUserHandler(userUC)
		// users := api.Group("/users")
		// {
		//     users.GET("/", userHandler.GetAll)
		//     users.POST("/", userHandler.Create)
		// }
	}
}
