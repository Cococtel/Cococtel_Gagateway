package catalogrepository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type (
	IAI interface {
		ProcessStrings(input []string) (string, error)
		CreateRecipe(liquor string) (*entities.AIRecipe, error)
		ExtractTextFromImage(imageBytes []byte) ([]string, error)
	}
	aiRepository struct{}
)

var ms_ai_endpoint string
var ms_ai_image_recognition string

func NewAIRepository() IAI {
	ms_ai_endpoint = os.Getenv("MS_AI_DOMAIN")
	ms_ai_image_recognition = os.Getenv("MS_IMAGE_RECOGNITION_DOMAIN")
	return &aiRepository{}
}

func (ir *aiRepository) ProcessStrings(input []string) (string, error) {
	url := fmt.Sprintf("%s/DeduceLiquorName", ms_ai_endpoint)
	body, _ := json.Marshal(input)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resultBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resultBytes), nil
}

func (ir *aiRepository) CreateRecipe(liquor string) (*entities.AIRecipe, error) {
	url := fmt.Sprintf("%s/CreateRecipe?liquor=%s", ms_ai_endpoint, liquor)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recipe entities.AIRecipe
	if err := json.NewDecoder(resp.Body).Decode(&recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (ir *aiRepository) ExtractTextFromImage(imageBytes []byte) ([]string, error) {
	// Crear un buffer y un multipart writer para construir la solicitud form-data.
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Crear el campo "imageFile" con un nombre de archivo (por ejemplo, "image.jpg").
	part, err := writer.CreateFormFile("imageFile", "image.jpg")
	if err != nil {
		return nil, err
	}
	// Escribir los bytes de la imagen en el campo.
	if _, err := part.Write(imageBytes); err != nil {
		return nil, err
	}
	// Es importante cerrar el writer para que se añada el boundary al Content-Type.
	if err := writer.Close(); err != nil {
		return nil, err
	}

	// Construir la solicitud HTTP con el body del multipart.
	req, err := http.NewRequest(http.MethodPost, ms_ai_image_recognition, &buf)
	if err != nil {
		return nil, err
	}
	// Establecer el Content-Type con el boundary generado.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.Status)
	}

	var texts []string
	if err := json.NewDecoder(resp.Body).Decode(&texts); err != nil {
		return nil, err
	}
	return texts, nil
}
