package catalogrepository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	IAI interface {
		ProcessStrings(input []string) (string, error)
		CreateRecipe(liquor string) (*entities.AIRecipe, error)
	}
	aiRepository struct{}
)

var ms_ai_endpoint string

func NewAIRepository() IAI {
	ms_ai_endpoint = os.Getenv("MS_AI_DOMAIN")
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
