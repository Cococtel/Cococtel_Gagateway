package postcontroller

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/postservice"
	"net/http"

	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	postsService postservice.PostsService
}

func NewPostsController(service postservice.PostsService) *PostsController {
	return &PostsController{postsService: service}
}

func (c *PostsController) GetPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, apiErr := c.postsService.GetPosts()
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": apiErr.Message().Error(),
					"status":  apiErr.Status(),
				},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  posts,
			"error": nil,
		})
	}
}

func (c *PostsController) GetPostByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		post, apiErr := c.postsService.GetPostByID(id)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": apiErr.Message().Error(),
					"status":  apiErr.Status(),
				},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  post,
			"error": nil,
		})
	}
}

func (c *PostsController) CreatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post entities.Post
		if err := ctx.ShouldBindJSON(&post); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": "invalid input",
					"status":  http.StatusBadRequest,
				},
			})
			return
		}
		newPost, apiErr := c.postsService.CreatePost(post)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": apiErr.Message().Error(),
					"status":  apiErr.Status(),
				},
			})
			return
		}
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{
			"data":  newPost,
			"error": nil,
		})
	}
}

func (c *PostsController) UpdatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updates map[string]interface{}
		if err := ctx.ShouldBindJSON(&updates); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": "invalid input",
					"status":  http.StatusBadRequest,
				},
			})
			return
		}
		updatedPost, apiErr := c.postsService.UpdatePost(id, updates)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": apiErr.Message().Error(),
					"status":  apiErr.Status(),
				},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  updatedPost,
			"error": nil,
		})
	}
}

func (c *PostsController) DeletePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		apiErr := c.postsService.DeletePost(id)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data": nil,
				"error": map[string]interface{}{
					"message": apiErr.Message().Error(),
					"status":  apiErr.Status(),
				},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  "post deleted successfully",
			"error": nil,
		})
	}
}
