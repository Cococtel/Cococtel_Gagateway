package main

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os"
	"strings"

	"github.com/Cococtel/Cococtel_Gagateway/internal/http"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var validAPIKeys []string

func main() {
	godotenv.Load(".env")
	err := getApiKeys()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	utils.InitMetrics()
	middleware.InitMiddlewareMetrics()

	eng := gin.Default()
	eng.Use(middleware.MetricsMiddleware())

	eng.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router := http.InitRouter(eng, nil)
	router.MapRoutes(validAPIKeys)

	log.Println("API Gateway corriendo en Cloud Run...")
	if err := eng.Run(); err != nil {
		panic(err)
	}
}

func getApiKeys() error {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in main.go: %s environment variable not set.", k)
		}
		return v
	}
	apiKeys := mustGetenv("VALID_API_KEYS")
	validAPIKeys = strings.Split(apiKeys, ",")
	if len(validAPIKeys) == 0 {
		return errors.New("No Valid API keys found")
	}
	return nil
}
