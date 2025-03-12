package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func InitUserHandler(e *gin.Engine, userService *services.UserService) {
	// Init handler
	h := &UserHandler{ UserService: userService }

	// Add routes to engine
	g := e.Group("api/v1/users")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getUserById)
		g.GET("/email/:email", h.getUserByEmail)
		g.GET("", h.getAllUsers)
		g.POST("", h.createUser)
		g.DELETE("/:id", h.deleteUserById)
		g.DELETE("/email/:email", h.deleteUserByEmail)
	}
}

func (h *UserHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.UserService.Health() })
}

func (h *UserHandler) getAllUsers(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetUserQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get users
	users, err := h.UserService.GetAllUsers(&query.Limit, &query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all users failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all users successfully", "data": users })
}

func (h *UserHandler) getUserById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetUserByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Get user
	user, err := h.UserService.GetUserById(&param.ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "User not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get user by id successfully", "data": user })
}

func (h *UserHandler) getUserByEmail(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetUserByEmailParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Get user
	user, err := h.UserService.GetUserByEmail(&param.Email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "User not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get user by email successfully", "data": user })
}

func (h *UserHandler) createUser(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateUserPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create user
	user, err := h.UserService.CreateUser(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create user failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Create user successfully", "data": user })
}

func (h *UserHandler) deleteUserById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetUserByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Delete user
	err := h.UserService.DeleteUserById(&param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete user by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete user by id successfully" })
}

func (h *UserHandler) deleteUserByEmail(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetUserByEmailParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Delete user
	err := h.UserService.DeleteUserByEmail(&param.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete user by email failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete user by email successfully" })
}
