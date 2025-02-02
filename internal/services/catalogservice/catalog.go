package catalogservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/catalogrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
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
		return nil, utils.NewApiError(errors.New("error getting liquors"), http.StatusInternalServerError)
	}
	return liquors, nil
}

func (cs *catalogService) GetLiquorByID(id string) (*entities.Liquor, utils.ApiError) {
	liquor, err := cs.catalogRepository.FetchLiquorByID(id)
	if err != nil {
		return nil, utils.NewApiError(errors.New("liquor not found"), http.StatusNotFound)
	}
	return liquor, nil
}

func (cs *catalogService) CreateLiquor(liquor dtos.Liquor) (*entities.Liquor, utils.ApiError) {
	newLiquor, err := cs.catalogRepository.CreateLiquor(liquor)
	if err != nil {
		return nil, utils.NewApiError(errors.New("error saving liquor"), http.StatusInternalServerError)
	}
	return newLiquor, nil
}

func (cs *catalogService) UpdateLiquor(id string, updates map[string]interface{}) (*entities.Liquor, utils.ApiError) {
	currentLiquor, apiErr := cs.catalogRepository.FetchLiquorByID(id)
	if apiErr != nil {
		return nil, utils.NewApiError(apiErr, http.StatusNotFound)
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
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	return updatedLiquor, nil
}

func (cs *catalogService) DeleteLiquor(id string) utils.ApiError {
	err := cs.catalogRepository.DeleteLiquor(id)
	if err != nil {
		return utils.NewApiError(errors.New("error deleting liquor"), http.StatusInternalServerError)
	}
	return nil
}

func (cs *catalogService) GetRecipes() ([]entities.Recipe, utils.ApiError) {
	recipes, err := cs.catalogRepository.FetchRecipes()
	if err != nil {
		return nil, utils.NewApiError(errors.New("error getting recipes"), http.StatusInternalServerError)
	}
	return recipes, nil
}

func (cs *catalogService) GetRecipeByID(id string) (*entities.Recipe, utils.ApiError) {
	recipe, err := cs.catalogRepository.FetchRecipeByID(id)
	if err != nil {
		return nil, utils.NewApiError(errors.New("recipe not found"), http.StatusNotFound)
	}
	return recipe, nil
}

func (cs *catalogService) CreateRecipe(recipe dtos.Recipe) (*entities.Recipe, utils.ApiError) {
	newRecipe, err := cs.catalogRepository.CreateRecipe(recipe)
	if err != nil {
		return nil, utils.NewApiError(errors.New("error saving recipe"), http.StatusInternalServerError)
	}
	return newRecipe, nil
}

func (cs *catalogService) UpdateRecipe(id string, updates map[string]interface{}) (*entities.Recipe, utils.ApiError) {
	currentRecipe, apiErr := cs.catalogRepository.FetchRecipeByID(id)
	if apiErr != nil {
		return nil, utils.NewApiError(apiErr, http.StatusNotFound)
	}

	updatedFields := make(map[string]interface{})

	if name, exists := updates["name"]; exists && name != currentRecipe.Name {
		updatedFields["name"] = name
	}
	if ingredients, exists := updates["ingredients"]; exists {
		updatedFields["ingredients"] = ingredients
	}
	if instructions, exists := updates["instructions"]; exists && instructions != currentRecipe.Instructions {
		updatedFields["instructions"] = instructions
	}
	if category, exists := updates["category"]; exists && category != currentRecipe.Category {
		updatedFields["category"] = category
	}
	if liquors, exists := updates["liquors"]; exists {
		updatedFields["liquors"] = liquors
	}

	if len(updatedFields) == 0 {
		return currentRecipe, nil
	}

	updatedRecipe, err := cs.catalogRepository.UpdateRecipe(id, updatedFields)
	if err != nil {
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	return updatedRecipe, nil
}

func (cs *catalogService) DeleteRecipe(id string) utils.ApiError {
	err := cs.catalogRepository.DeleteRecipe(id)
	if err != nil {
		return utils.NewApiError(errors.New("error al eliminar receta"), http.StatusInternalServerError)
	}
	return nil
}
