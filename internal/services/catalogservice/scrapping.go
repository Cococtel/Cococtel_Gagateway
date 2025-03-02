package catalogservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/catalogrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"log"
)

type (
	IScrapping interface {
		GetProductByCode(code string) (*entities.Product, utils.ApiError)
	}
	scrappingService struct {
		scrappingRepository catalogrepository.IScrapping
	}
)

func NewScrappingService(repo catalogrepository.IScrapping) IScrapping {
	return &scrappingService{scrappingRepository: repo}
}

func (ss *scrappingService) GetProductByCode(code string) (*entities.Product, utils.ApiError) {
	product, err := ss.scrappingRepository.GetProductByCode(code)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("product not found"), err.Status())
	}
	return product, nil
}
