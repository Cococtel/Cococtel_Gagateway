package catalogcontroller

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAI interface {
		ProcessStrings() gin.HandlerFunc
		CreateRecipe() gin.HandlerFunc
	}
	aiController struct {
		aiService catalogservice.IAI
	}
)

func NewAIController(service catalogservice.IAI) *aiController {
	return &aiController{aiService: service}
}

func (ai *aiController) ProcessStrings() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input []string
		if err := ctx.ShouldBindJSON(&input); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "invalid input", "status": http.StatusBadRequest},
			})
			return
		}

		result, apiErr := ai.aiService.ProcessStrings(input)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  result,
			"error": nil,
		})
	}
}

func (ai *aiController) CreateRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		liquor := ctx.Query("liquor")
		if liquor == "" {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "liquor required", "status": http.StatusBadRequest},
			})
			return
		}

		recipe, apiErr := ai.aiService.CreateRecipe(liquor)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  recipe,
			"error": nil,
		})
	}
}
