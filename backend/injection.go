package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nambuitechx/go-social/configs"
	"github.com/nambuitechx/go-social/handlers"
	"github.com/nambuitechx/go-social/repositories"
	"github.com/nambuitechx/go-social/services"
)

func getEngine() *gin.Engine {
	// Connect database
	settings := configs.NewSettings()
	db := configs.NewDatabaseConnection(settings).DB

	// Repositories
	userRepository := repositories.NewUserRepository(db)
	postRepository := repositories.NewPostRepository(db)

	// Services
	userService := services.NewUserService(userRepository)	
	postService := services.NewPostService(postRepository)	

	// Routes
	engine := gin.Default()

	engine.GET("/health", checkHealth)
	handlers.InitUserHandler(&handlers.HandlerConfig{ Engine: engine }, userService)
	handlers.InitPostHandler(&handlers.HandlerConfig{ Engine: engine }, postService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
