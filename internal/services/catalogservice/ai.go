package catalogservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/catalogrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"net/http"
)

type (
	IAI interface {
		ProcessStrings(input []string) (string, utils.ApiError)
		CreateRecipe(liquor string) (*entities.AIRecipe, utils.ApiError)
		ExtractTextFromImage(imageBytes []byte) ([]string, utils.ApiError)
	}
	aiService struct {
		aiRepository catalogrepository.IAI
	}
)

func NewAIService(repo catalogrepository.IAI) IAI {
	return &aiService{aiRepository: repo}
}

func (is *aiService) ProcessStrings(input []string) (string, utils.ApiError) {
	result, err := is.aiRepository.ProcessStrings(input)
	if err != nil {
		return "", utils.NewApiError(errors.New("error getting liquor"), http.StatusInternalServerError)
	}
	return result, nil
}

func (is *aiService) CreateRecipe(liquor string) (*entities.AIRecipe, utils.ApiError) {
	recipe, err := is.aiRepository.CreateRecipe(liquor)
	if err != nil {
		return nil, utils.NewApiError(errors.New("error generating recipe"), http.StatusInternalServerError)
	}
	return recipe, nil
}

func (is *aiService) ExtractTextFromImage(imageBytes []byte) ([]string, utils.ApiError) {
	texts, err := is.aiRepository.ExtractTextFromImage(imageBytes)
	if err != nil {
		return nil, utils.NewApiError(errors.New("error extracting text from image"), http.StatusInternalServerError)
	}
	return texts, nil
}
