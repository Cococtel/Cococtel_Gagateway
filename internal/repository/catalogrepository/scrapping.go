package catalogrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"net/http"
	"os"
)

type (
	IScrapping interface {
		GetProductByCode(code string) (*entities.Product, error)
	}
	scrappingRepository struct{}
)

var ms_scrapping_endpoint string

func NewScrappingRepository() IScrapping {
	ms_scrapping_endpoint = os.Getenv("MS_SCRAPPING_DOMAIN")
	return &scrappingRepository{}
}

func (sr *scrappingRepository) GetProductByCode(code string) (*entities.Product, error) {
	url := fmt.Sprintf("%s/%s", ms_scrapping_endpoint, code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var product entities.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	if product.Name == "" {
		return nil, errors.New("product not found")
	}

	return &product, nil
}
