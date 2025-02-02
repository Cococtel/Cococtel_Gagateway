package catalogcontroller

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	ICatalog interface {
		GetLiquors() gin.HandlerFunc
		GetLiquorByID() gin.HandlerFunc
		CreateLiquor() gin.HandlerFunc
		UpdateLiquor() gin.HandlerFunc
		DeleteLiquor() gin.HandlerFunc
		GetRecipes() gin.HandlerFunc
		GetRecipeByID() gin.HandlerFunc
		CreateRecipe() gin.HandlerFunc
		UpdateRecipe() gin.HandlerFunc
		DeleteRecipe() gin.HandlerFunc
	}

	catalogController struct {
		catalogService catalogservice.ICatalog
	}
)

func NewLiquorController(service catalogservice.ICatalog) ICatalog {
	return &catalogController{catalogService: service}
}

func (c *catalogController) GetLiquors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		liquors, apiErr := c.catalogService.GetLiquors()

		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  liquors,
			"error": nil,
		})
	}
}

func (c *catalogController) GetLiquorByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		liquor, apiErr := c.catalogService.GetLiquorByID(id)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  liquor,
			"error": nil,
		})
	}
}

func (c *catalogController) CreateLiquor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var liquor dtos.Liquor
		if err := ctx.ShouldBindJSON(&liquor); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "Datos inválidos", "status": http.StatusBadRequest},
			})
			return
		}

		newLiquor, apiErr := c.catalogService.CreateLiquor(liquor)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusCreated, map[string]interface{}{
			"data":  newLiquor,
			"error": nil,
		})
	}
}

func (c *catalogController) UpdateLiquor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updates map[string]interface{}
		if err := ctx.ShouldBindJSON(&updates); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "Datos inválidos", "status": http.StatusBadRequest},
			})
			return
		}

		updatedLiquor, apiErr := c.catalogService.UpdateLiquor(id, updates)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  updatedLiquor,
			"error": nil,
		})
	}
}

func (c *catalogController) DeleteLiquor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		apiErr := c.catalogService.DeleteLiquor(id)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  "liquor deleted successfully",
			"error": nil,
		})
	}
}

func (c *catalogController) GetRecipes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recipes, apiErr := c.catalogService.GetRecipes()
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  recipes,
			"error": nil,
		})
	}
}

func (c *catalogController) GetRecipeByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		recipe, apiErr := c.catalogService.GetRecipeByID(id)
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

func (c *catalogController) CreateRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var recipe dtos.Recipe
		if err := ctx.ShouldBindJSON(&recipe); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "Datos inválidos", "status": http.StatusBadRequest},
			})
			return
		}

		newRecipe, apiErr := c.catalogService.CreateRecipe(recipe)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusCreated, map[string]interface{}{
			"data":  newRecipe,
			"error": nil,
		})
	}
}

func (c *catalogController) UpdateRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updates map[string]interface{}

		if err := ctx.ShouldBindJSON(&updates); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "Datos inválidos", "status": http.StatusBadRequest},
			})
			return
		}

		updatedRecipe, apiErr := c.catalogService.UpdateRecipe(id, updates)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  updatedRecipe,
			"error": nil,
		})
	}
}

func (c *catalogController) DeleteRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		apiErr := c.catalogService.DeleteRecipe(id)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  "Recipe deleted successfully",
			"error": nil,
		})
	}
}
