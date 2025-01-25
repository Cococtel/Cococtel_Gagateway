package middleware

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/defines"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateAPIKey(apiKeys []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey == "" {
			utils.Error(c, http.StatusUnauthorized, defines.InvalidApiKey.Error())
			c.Abort()
			return
		}

		isValid := false
		for _, key := range apiKeys {
			if key == apiKey {
				isValid = true
				break
			}
		}

		if !isValid {
			utils.Error(c, http.StatusUnauthorized, defines.InvalidApiKey.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
