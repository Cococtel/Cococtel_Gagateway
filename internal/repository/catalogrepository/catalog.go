package catalogrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
)

type (
	ICatalog interface {
		FetchLiquors() ([]entities.Liquor, utils.ApiError)
		FetchLiquorByID(id string) (*entities.Liquor, utils.ApiError)
		CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, utils.ApiError)
		UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, utils.ApiError)
		DeleteLiquor(id string) utils.ApiError
		FetchRecipes() ([]entities.Recipe, utils.ApiError)
		FetchRecipeByID(id string) (*entities.Recipe, utils.ApiError)
		CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, utils.ApiError)
		UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, utils.ApiError)
		DeleteRecipe(id string) utils.ApiError
	}
	catalogRepository struct{}
)

var ms_catalog_endpoint string

func NewCatalogRepository() ICatalog {
	ms_catalog_endpoint = os.Getenv("MS_CATALOG_DOMAIN")
	return &catalogRepository{}
}

func (cr *catalogRepository) FetchLiquors() ([]entities.Liquor, utils.ApiError) {
	start := time.Now()
	resp, err := http.Get(ms_catalog_endpoint + "/liquors")
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "FetchLiquors", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var liquors []entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&liquors); err != nil {
		utils.MeasureRequest("catalog", "FetchLiquors", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "FetchLiquors", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("error getting liquors"), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "FetchLiquors", start, resp.StatusCode, nil)
	return liquors, nil
}

func (cr *catalogRepository) FetchLiquorByID(id string) (*entities.Liquor, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "FetchLiquorByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var liquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&liquor); err != nil {
		utils.MeasureRequest("catalog", "FetchLiquorByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "FetchLiquorByID", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("error getting liquor"), resp.StatusCode)
	}
	utils.MeasureRequest("catalog", "FetchLiquorByID", start, resp.StatusCode, nil)
	return &liquor, nil
}

func (cr *catalogRepository) CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, utils.ApiError) {
	start := time.Now()
	body, _ := json.Marshal(liquor)
	resp, err := http.Post(ms_catalog_endpoint+"/liquors", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "CreateLiquor", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var newLiquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&newLiquor); err != nil {
		utils.MeasureRequest("catalog", "CreateLiquor", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "CreateLiquor", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "CreateLiquor", start, resp.StatusCode, nil)
	return &newLiquor, nil
}

func (cr *catalogRepository) UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)
	body, _ := json.Marshal(updates)

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "UpdateLiquor", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var updatedLiquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&updatedLiquor); err != nil {
		utils.MeasureRequest("catalog", "UpdateLiquor", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "UpdateLiquor", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "UpdateLiquor", start, resp.StatusCode, nil)
	return &updatedLiquor, nil
}

func (cr *catalogRepository) DeleteLiquor(id string) utils.ApiError {
	start := time.Now()
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "DeleteLiquor", start, resp.StatusCode, err)
		return utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "DeleteLiquor", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}
	utils.MeasureRequest("catalog", "DeleteLiquor", start, resp.StatusCode, nil)
	return nil
}

func (cr *catalogRepository) FetchRecipes() ([]entities.Recipe, utils.ApiError) {
	start := time.Now()
	resp, err := http.Get(ms_catalog_endpoint + "/recipes")
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "FetchRecipes", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var recipes []entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipes); err != nil {
		utils.MeasureRequest("catalog", "FetchRecipes", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "FetchRecipes", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "FetchRecipes", start, resp.StatusCode, nil)
	return recipes, nil
}

func (cr *catalogRepository) FetchRecipeByID(id string) (*entities.Recipe, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "FetchRecipeByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var recipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipe); err != nil {
		utils.MeasureRequest("catalog", "FetchRecipeByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "FetchRecipeByID", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "FetchRecipeByID", start, resp.StatusCode, nil)
	return &recipe, nil
}

func (cr *catalogRepository) CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, utils.ApiError) {
	start := time.Now()
	body, _ := json.Marshal(recipe)
	resp, err := http.Post(ms_catalog_endpoint+"/recipes", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "CreateRecipe", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var newRecipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&newRecipe); err != nil {
		utils.MeasureRequest("catalog", "CreateRecipe", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "CreateRecipe", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	utils.MeasureRequest("catalog", "CreateRecipe", start, resp.StatusCode, nil)
	return &newRecipe, nil
}

func (cr *catalogRepository) UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)

	// Convertir updates a JSON
	body, err := json.Marshal(updates)
	if err != nil {
		log.Println("Error marshalling update data:", err)
		utils.MeasureRequest("catalog", "UpdateRecipe", start, http.StatusInternalServerError, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		utils.MeasureRequest("catalog", "UpdateRecipe", start, http.StatusInternalServerError, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request failed:", err)
		utils.MeasureRequest("catalog", "UpdateRecipe", start, http.StatusInternalServerError, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Manejo de errores HTTP
	if resp.StatusCode >= http.StatusMultipleChoices {
		log.Printf("UpdateRecipe failed with status %d\n", resp.StatusCode)
		utils.MeasureRequest("catalog", "UpdateRecipe", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}

	var updatedRecipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&updatedRecipe); err != nil {
		log.Println("Error decoding response:", err)
		utils.MeasureRequest("catalog", "UpdateRecipe", start, http.StatusInternalServerError, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	utils.MeasureRequest("catalog", "UpdateRecipe", start, resp.StatusCode, nil)
	return &updatedRecipe, nil
}

func (cr *catalogRepository) DeleteRecipe(id string) utils.ApiError {
	start := time.Now()
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("catalog", "DeleteRecipe", start, resp.StatusCode, err)
		return utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("catalog", "DeleteRecipe", start, resp.StatusCode, errors.New(http.StatusText(resp.StatusCode)))
		return utils.NewApiError(errors.New(http.StatusText(resp.StatusCode)), resp.StatusCode)
	}
	utils.MeasureRequest("catalog", "DeleteRecipe", start, resp.StatusCode, nil)
	return nil
}
