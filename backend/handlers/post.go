package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/services"
)

type PostHandler struct {
	PostService *services.PostService
}

func InitPostHandler(c *HandlerConfig, postService *services.PostService) {
	// Init handler
	h := &PostHandler{ PostService: postService }

	// Add routes to engine
	g := c.Engine.Group("api/v1/posts")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getPostById)
		g.GET("", h.getAllPosts)
		g.POST("", h.createPost)
		g.DELETE("/:id", h.deletePostById)
	}
}

func (h *PostHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.PostService.Health() })
}

func (h *PostHandler) getAllPosts(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetPostQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get posts
	posts := h.PostService.GetAllPosts(&query.Limit, &query.Offset)

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all posts successfully", "data": posts })
}

func (h *PostHandler) getPostById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetPostByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Get post
	post, err := h.PostService.GetPostById(&param.ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Post not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get post by id successfully", "data": post })
}

func (h *PostHandler) createPost(ctx *gin.Context) {
	// Get payload
	userId := "2ad2d0c7-1965-4109-8565-69407273b28c" // email: nambui - for test without auth
	payload := &models.CreatePostPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create post
	post, err := h.PostService.CreatePost(payload, userId);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create post failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Create post successfully", "data": post })
}

func (h *PostHandler) deletePostById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetPostByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Delete post
	err := h.PostService.DeletePostById(&param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete post by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete post by id successfully" })
}
