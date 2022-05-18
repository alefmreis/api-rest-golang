package services

import (
	"errors"

	"github.com/alefmreis/api-rest-golang/src/core/domain"
	repositories "github.com/alefmreis/api-rest-golang/src/core/ports/repositories"
	"github.com/alefmreis/api-rest-golang/src/helpers"
)

type BrandService struct {
	BrandRepository repositories.BrandRepository
}

func NewBrandService(brandRepository repositories.BrandRepository) *BrandService {
	return &BrandService{
		BrandRepository: brandRepository,
	}
}

func (cbuc *BrandService) Save(name string) error {

	brand := domain.NewBrand(helpers.NewUUID().New(), name)

	if err := cbuc.BrandRepository.Create(brand); err != nil {
		return errors.New("create brand into repository has failed")
	}

	return nil
}
