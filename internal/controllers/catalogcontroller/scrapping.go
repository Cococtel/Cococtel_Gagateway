package catalogcontroller

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IScrapping interface {
		GetProductByCode() gin.HandlerFunc
	}
	scrappingController struct {
		aiService catalogservice.IScrapping
	}
)

func NewScrappingController(service catalogservice.IScrapping) *scrappingController {
	return &scrappingController{aiService: service}
}

func (s *scrappingController) GetProductByCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")

		product, apiErr := s.aiService.GetProductByCode(code)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  product,
			"error": nil,
		})
	}
}
