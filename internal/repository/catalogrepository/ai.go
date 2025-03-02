package catalogrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
)

type (
	IAI interface {
		ProcessStrings(input []string) (string, utils.ApiError)
		CreateRecipe(liquor string) (*entities.AIRecipe, utils.ApiError)
		ExtractTextFromImage(imageBytes []byte) ([]string, utils.ApiError)
	}
	aiRepository struct{}
)

var (
	msAIEndpoint         string
	msAIImageRecognition string
)

func NewAIRepository() IAI {
	msAIEndpoint = os.Getenv("MS_AI_DOMAIN")
	msAIImageRecognition = os.Getenv("MS_IMAGE_RECOGNITION_DOMAIN")
	return &aiRepository{}
}

func (ir *aiRepository) ProcessStrings(input []string) (string, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/DeduceLiquorName", msAIEndpoint)

	body, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ProcessStrings", start, http.StatusBadRequest, err)
		return "", utils.NewApiError(err, http.StatusBadRequest)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ProcessStrings", start, http.StatusInternalServerError, err)
		return "", utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	resultBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ProcessStrings", start, resp.StatusCode, err)
		return "", utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(string(resultBytes))
		utils.MeasureRequest("ai", "ProcessStrings", start, resp.StatusCode, errors.New(resp.Status))
		return "", utils.NewApiError(err, resp.StatusCode)
	}

	utils.MeasureRequest("ai", "ProcessStrings", start, resp.StatusCode, nil)
	return string(resultBytes), nil
}

func (ir *aiRepository) CreateRecipe(liquor string) (*entities.AIRecipe, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/CreateRecipe?liquor=%s", msAIEndpoint, liquor)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "CreateRecipe", start, http.StatusInternalServerError, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var recipe entities.AIRecipe
	if err := json.NewDecoder(resp.Body).Decode(&recipe); err != nil {
		utils.MeasureRequest("ai", "CreateRecipe", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("ai", "CreateRecipe", start, resp.StatusCode, errors.New(resp.Status))
		return nil, utils.NewApiError(errors.New("re"), resp.StatusCode)
	}

	utils.MeasureRequest("ai", "CreateRecipe", start, resp.StatusCode, nil)
	return &recipe, nil
}

func (ir *aiRepository) ExtractTextFromImage(imageBytes []byte) ([]string, utils.ApiError) {
	start := time.Now()
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("imageFile", "image.jpg")
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	if _, err := part.Write(imageBytes); err != nil {
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	if err := writer.Close(); err != nil {
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	req, err := http.NewRequest(http.MethodPost, msAIImageRecognition, &buf)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, 500, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices {
		err := fmt.Errorf("status error: %v", resp.Status)
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, resp.StatusCode, errors.New(resp.Status))
		return nil, utils.NewApiError(err, resp.StatusCode)
	}

	var texts []string
	if err := json.NewDecoder(resp.Body).Decode(&texts); err != nil {
		utils.MeasureRequest("ai", "ExtractTextFromImage", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	utils.MeasureRequest("ai", "ExtractTextFromImage", start, resp.StatusCode, nil)
	return texts, nil
}
