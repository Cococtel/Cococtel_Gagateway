package main

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var validAPIKeys []string

func main() {
	godotenv.Load(".env")
	err := getApiKeys()
	if err != nil {
		panic(err)
	}

	eng := gin.Default()
	router := http.InitRouter(eng, nil)
	router.MapRoutes(validAPIKeys)
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
