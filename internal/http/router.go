package http

import (
	"database/sql"
	"github.com/Cococtel/Cococtel_Gagateway/internal/controllers"
	"github.com/Cococtel/Cococtel_Gagateway/internal/controllers/authcontroller"
	"github.com/Cococtel/Cococtel_Gagateway/internal/controllers/catalogcontroller"
	"github.com/Cococtel/Cococtel_Gagateway/internal/defines"
	"github.com/Cococtel/Cococtel_Gagateway/internal/graph"
	"github.com/Cococtel/Cococtel_Gagateway/internal/middleware"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/authrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/catalogrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/authservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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
	catalogRepository := catalogrepository.NewCatalogRepository()
	aiRepository := catalogrepository.NewAIRepository()
	scrappingRepository := catalogrepository.NewScrappingRepository()
	authRepository := authrepository.NewAuthRepository()

	catalogService := catalogservice.NewCatalogService(catalogRepository)
	aiService := catalogservice.NewAIService(aiRepository)
	scrappingService := catalogservice.NewScrappingService(scrappingRepository)
	authService := authservice.NewAuthService(authRepository)

	catalogController := catalogcontroller.NewLiquorController(catalogService)
	aiController := catalogcontroller.NewAIController(aiService)
	scrappingController := catalogcontroller.NewScrappingController(scrappingService)
	authController := authcontroller.NewAuthController(authService)

	// REST Licores
	r.eng.GET("/liquors", catalogController.GetLiquors())
	r.eng.GET("/liquors/:id", catalogController.GetLiquorByID())
	r.eng.POST("/liquors", catalogController.CreateLiquor())
	r.eng.PUT("/liquors/:id", catalogController.UpdateLiquor())
	r.eng.DELETE("/liquors/:id", catalogController.DeleteLiquor())

	// REST Recetas
	r.eng.GET("/recipes", catalogController.GetRecipes())
	r.eng.GET("/recipes/:id", catalogController.GetRecipeByID())
	r.eng.POST("/recipes", catalogController.CreateRecipe())
	r.eng.PUT("/recipes/:id", catalogController.UpdateRecipe())
	r.eng.DELETE("/recipes/:id", catalogController.DeleteRecipe())

	// REST AI & Scrapping
	r.eng.POST("/processStrings", aiController.ProcessStrings())
	r.eng.POST("/createAIRecipe", aiController.CreateRecipe())
	r.eng.GET("/product/:code", scrappingController.GetProductByCode())

	// REST Auth
	r.eng.GET("/verify", authController.Verify())
	r.eng.POST("/register", authController.Register())
	r.eng.POST("/login", authController.Login())

	// GraphQL Config
	schema, err := graphql.NewSchema(graph.NewSchema(catalogService, authService, scrappingService, aiService))
	if err != nil {
		panic(err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.eng.GET("/graphql", gin.WrapH(h))
	r.eng.POST("/graphql", gin.WrapH(h))
}
func (r *router) addSystemPaths() {
	r.eng.GET(defines.PingPath, controllers.Ping())
}
