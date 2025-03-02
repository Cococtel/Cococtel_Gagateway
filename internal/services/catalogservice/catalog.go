package catalogservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/catalogrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"log"
	"net/http"
)

type (
	ICatalog interface {
		GetLiquors() ([]entities.Liquor, utils.ApiError)
		GetLiquorByID(id string) (*entities.Liquor, utils.ApiError)
		CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, utils.ApiError)
		UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, utils.ApiError)
		DeleteLiquor(id string) utils.ApiError
		GetRecipes() ([]entities.Recipe, utils.ApiError)
		GetRecipeByID(id string) (*entities.Recipe, utils.ApiError)
		CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, utils.ApiError)
		UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, utils.ApiError)
		DeleteRecipe(id string) utils.ApiError
	}

	catalogService struct {
		catalogRepository catalogrepository.ICatalog
	}
)

func NewCatalogService(repo catalogrepository.ICatalog) ICatalog {
	return &catalogService{catalogRepository: repo}
}

func (cs *catalogService) GetLiquors() ([]entities.Liquor, utils.ApiError) {
	liquors, err := cs.catalogRepository.FetchLiquors()
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error getting liquors"), err.Status())
	}
	return liquors, nil
}

func (cs *catalogService) GetLiquorByID(id string) (*entities.Liquor, utils.ApiError) {
	liquor, err := cs.catalogRepository.FetchLiquorByID(id)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("liquor not found"), err.Status())
	}
	return liquor, nil
}

func (cs *catalogService) CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, utils.ApiError) {
	newLiquor, err := cs.catalogRepository.CreateLiquor(liquor)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error saving liquor"), err.Status())
	}
	return newLiquor, nil
}

func (cs *catalogService) UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, utils.ApiError) {
	currentLiquor, apiErr := cs.catalogRepository.FetchLiquorByID(id)
	if apiErr != nil {
		return nil, apiErr
	}
	updatedFields := make(map[string]interface{})

	if name, exists := updates["name"]; exists && name != currentLiquor.Name {
		updatedFields["name"] = name
	}
	if EAN, exists := updates["EAN"]; exists && EAN != currentLiquor.EAN {
		updatedFields["EAN"] = EAN
	}
	if category, exists := updates["category"]; exists && category != currentLiquor.Category {
		updatedFields["category"] = category
	}
	if description, exists := updates["description"]; exists && description != currentLiquor.Description {
		updatedFields["description"] = description
	}
	if additionalAttributes, exists := updates["additional_attributes"]; exists && additionalAttributes != currentLiquor.AdditionalAttributes {
		updatedFields["additional_attributes"] = additionalAttributes
	}

	if len(updatedFields) == 0 {
		return currentLiquor, nil
	}

	updatedLiquor, err := cs.catalogRepository.UpdateLiquor(id, updatedFields)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return updatedLiquor, nil
}

func (cs *catalogService) DeleteLiquor(id string) utils.ApiError {
	err := cs.catalogRepository.DeleteLiquor(id)
	if err != nil {
		log.Println(err)
		return utils.NewApiError(errors.New("error deleting liquor"), err.Status())
	}
	return nil
}

func (cs *catalogService) GetRecipes() ([]entities.Recipe, utils.ApiError) {
	recipes, err := cs.catalogRepository.FetchRecipes()
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error getting recipes"), err.Status())
	}
	return recipes, nil
}

func (cs *catalogService) GetRecipeByID(id string) (*entities.Recipe, utils.ApiError) {
	recipe, err := cs.catalogRepository.FetchRecipeByID(id)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("recipe not found"), err.Status())
	}
	return recipe, nil
}

func (cs *catalogService) CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, utils.ApiError) {
	newRecipe, err := cs.catalogRepository.CreateRecipe(recipe)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error saving recipe"), err.Status())
	}
	return newRecipe, nil
}

func (cs *catalogService) UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, utils.ApiError) {
	if len(updates) == 0 {
		return nil, utils.NewApiError(errors.New("no fields to update"), http.StatusBadRequest)
	}

	updatedRecipe, err := cs.catalogRepository.UpdateRecipe(id, updates)
	if err != nil {
		log.Println("Error updating recipe:", err)
		return nil, utils.NewApiError(errors.New("error updating recipe"), err.Status())
	}
	return updatedRecipe, nil
}

func (cs *catalogService) DeleteRecipe(id string) utils.ApiError {
	err := cs.catalogRepository.DeleteRecipe(id)
	if err != nil {
		log.Println(err)
		return utils.NewApiError(errors.New("error deleting recipe"), err.Status())
	}
	return nil
}
