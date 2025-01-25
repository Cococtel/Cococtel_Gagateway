package http

import (
	"database/sql"
	"github.com/Cococtel/Cococtel_Gagateway/internal/controllers"
	"github.com/Cococtel/Cococtel_Gagateway/internal/defines"
	"github.com/Cococtel/Cococtel_Gagateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes([]string)
}

type router struct {
	eng *gin.Engine
	db  *sql.DB
}

func (r *router) MapRoutes(validAPIKeys []string) {
	r.setGroup(validAPIKeys)
	r.addSystemPaths()
	r.buildRoutes()
}

func (r *router) setGroup(apiKeys []string) {
	r.eng.Use(middleware.CORS(), middleware.ValidateAPIKey(apiKeys))
}

func (r *router) buildRoutes() {
}
func (r *router) addSystemPaths() {
	r.eng.GET(defines.PingPath, controllers.Ping())
}
