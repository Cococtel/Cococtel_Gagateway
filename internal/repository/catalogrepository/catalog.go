package catalogrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"net/http"
	"os"
)

type (
	ICatalog interface {
		FetchLiquors() ([]entities.Liquor, error)
		FetchLiquorByID(id string) (*entities.Liquor, error)
		CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, error)
		UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, error)
		DeleteLiquor(id string) error
		FetchRecipes() ([]entities.Recipe, error)
		FetchRecipeByID(id string) (*entities.Recipe, error)
		CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, error)
		UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, error)
		DeleteRecipe(id string) error
	}
	catalogRepository struct{}
)

var ms_catalog_endpoint string

func NewCatalogRepository() ICatalog {
	ms_catalog_endpoint = os.Getenv("MS_CATALOG_DOMAIN")
	return &catalogRepository{}
}

func (cr *catalogRepository) FetchLiquors() ([]entities.Liquor, error) {
	resp, err := http.Get(ms_catalog_endpoint + "/liquors")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var liquors []entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&liquors); err != nil {
		return nil, err
	}

	if len(liquors) == 0 {
		return nil, errors.New("liquors not found")
	}

	return liquors, nil
}

func (cr *catalogRepository) FetchLiquorByID(id string) (*entities.Liquor, error) {
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var liquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&liquor); err != nil {
		return nil, err
	}

	if liquor.ID == "" {
		return nil, errors.New("liquor not found")
	}

	return &liquor, nil
}

func (cr *catalogRepository) CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, error) {
	body, _ := json.Marshal(liquor)
	resp, err := http.Post(ms_catalog_endpoint+"/liquors", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var newLiquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&newLiquor); err != nil {
		return nil, err
	}

	return &newLiquor, nil
}

func (cr *catalogRepository) UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, error) {
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)
	body, _ := json.Marshal(updates)

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedLiquor entities.Liquor
	if err := json.NewDecoder(resp.Body).Decode(&updatedLiquor); err != nil {
		return nil, err
	}

	return &updatedLiquor, nil
}

func (cr *catalogRepository) DeleteLiquor(id string) error {
	url := fmt.Sprintf("%s/liquors/%s", ms_catalog_endpoint, id)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (cr *catalogRepository) FetchRecipes() ([]entities.Recipe, error) {
	resp, err := http.Get(ms_catalog_endpoint + "/recipes")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recipes []entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (cr *catalogRepository) FetchRecipeByID(id string) (*entities.Recipe, error) {
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (cr *catalogRepository) CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, error) {
	body, _ := json.Marshal(recipe)
	resp, err := http.Post(ms_catalog_endpoint+"/recipes", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var newRecipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&newRecipe); err != nil {
		return nil, err
	}

	return &newRecipe, nil
}

func (cr *catalogRepository) UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, error) {
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)
	body, _ := json.Marshal(updates)

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedRecipe entities.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&updatedRecipe); err != nil {
		return nil, err
	}

	return &updatedRecipe, nil
}

func (cr *catalogRepository) DeleteRecipe(id string) error {
	url := fmt.Sprintf("%s/recipes/%s", ms_catalog_endpoint, id)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
