package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/nambuitechx/go-social/configs"
	"github.com/nambuitechx/go-social/handlers"
	"github.com/nambuitechx/go-social/repositories"
	"github.com/nambuitechx/go-social/services"
)

func getEngine() *gin.Engine {
	// Config settings
	settings := configs.NewSettings()

	// Connect database
	db := configs.NewDatabaseConnection(settings)

	// Repositories
	userRepository := repositories.NewUserRepository(db.DB)
	postRepository := repositories.NewPostRepository(db.DB)

	// Services
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)	
	postService := services.NewPostService(postRepository)

	// Engine
	engine := gin.Default()

	// Middlewares
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true

	engine.Use(cors.New(config))

	engine.GET("/health", checkHealth)
	handlers.InitAuthHandler(engine, authService)
	handlers.InitUserHandler(engine, userService)
	handlers.InitPostHandler(engine, postService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
