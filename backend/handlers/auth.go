package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func InitAuthHandler(e *gin.Engine, authService *services.AuthService) {
	// Init handler
	h := &AuthHandler{ AuthService: authService }

	// Add routes to engine
	g := e.Group("api/v1/auth")
	{
		g.GET("/health", h.health)
		g.POST("/register", h.register)
		g.POST("/login", h.login)
	}
}

func (h *AuthHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.AuthService.Health() })
}

func (h *AuthHandler) register(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateUserPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Register
	user, err := h.AuthService.Register(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Register failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Register successfully", "data": user })
}

func (h *AuthHandler) login(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateUserPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Login
	tokenInfo, err := h.AuthService.Login(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Login failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Login successfully", "data": tokenInfo })
}
