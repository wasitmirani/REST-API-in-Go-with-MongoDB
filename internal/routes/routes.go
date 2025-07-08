package routes

import (
	"crud-with-mongodb/internal/controllers"
	"crud-with-mongodb/internal/repositories"
	"crud-with-mongodb/internal/services"
	"crud-with-mongodb/pkg/middleware"
	"crud-with-mongodb/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	
)

func SetupRouter(db *mongo.Client, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db.Database("crudapp_db"))

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)

	// API routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", userController.CreateUser)
			users.GET("/", userController.GetAllUsers)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}