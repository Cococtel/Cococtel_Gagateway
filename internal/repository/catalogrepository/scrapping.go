package catalogrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
)

type (
	IScrapping interface {
		GetProductByCode(code string) (*entities.Product, utils.ApiError)
	}
	scrappingRepository struct{}
)

var ms_scrapping_endpoint string

func NewScrappingRepository() IScrapping {
	ms_scrapping_endpoint = os.Getenv("MS_SCRAPPING_DOMAIN")
	return &scrappingRepository{}
}

func (sr *scrappingRepository) GetProductByCode(code string) (*entities.Product, utils.ApiError) {
	start := time.Now() // Iniciar medición de latencia
	url := fmt.Sprintf("%s/%s", ms_scrapping_endpoint, code)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("scrapping", "GetProductByCode", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var product entities.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		utils.MeasureRequest("scrapping", "GetProductByCode", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	if resp.StatusCode > http.StatusMultipleChoices {
		utils.MeasureRequest("scrapping", "GetProductByCode", start, resp.StatusCode, errors.New("product not found"))
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	if product.Name == "" {
		err := errors.New("product not found")
		utils.MeasureRequest("scrapping", "GetProductByCode", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusNotFound)
	}

	utils.MeasureRequest("scrapping", "GetProductByCode", start, resp.StatusCode, nil)
	return &product, nil
}
