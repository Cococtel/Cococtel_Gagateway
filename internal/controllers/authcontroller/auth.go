package authcontroller

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/authservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAuth interface {
		Verify() gin.HandlerFunc
		Register() gin.HandlerFunc
		Login() gin.HandlerFunc
	}
	authController struct {
		authService authservice.IAuth
	}
)

func NewAuthController(authservice authservice.IAuth) IAuth {
	return &authController{authService: authservice}
}

func (a *authController) Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("x-auth-token")
		if token == "" {
			utils.Response(ctx, http.StatusUnauthorized, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "invalid x-auth-token", "status": http.StatusUnauthorized},
			})
			return
		}

		apiErr := a.authService.Verify(token)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  "successfully verify",
			"error": nil,
		})
	}
}

func (a *authController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user dtos.Register
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "invalid data", "status": http.StatusBadRequest},
			})
			return
		}

		newUser, apiErr := a.authService.Register(user)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusCreated, map[string]interface{}{
			"data":  newUser,
			"error": nil,
		})
	}
}

func (a *authController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentails dtos.Login
		if err := ctx.ShouldBindJSON(&credentails); err != nil {
			utils.Response(ctx, http.StatusBadRequest, map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": "invalid data", "status": http.StatusBadRequest},
			})
			return
		}

		loginResponse, apiErr := a.authService.Login(credentails)
		if apiErr != nil {
			utils.Response(ctx, apiErr.Status(), map[string]interface{}{
				"data":  nil,
				"error": map[string]interface{}{"message": apiErr.Message().Error(), "status": apiErr.Status()},
			})
			return
		}

		utils.Response(ctx, http.StatusOK, map[string]interface{}{
			"data":  loginResponse,
			"error": nil,
		})
	}
}
