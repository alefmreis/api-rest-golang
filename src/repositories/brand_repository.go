package repositories

import (
	"github.com/alefmreis/api-rest-golang/src/core/domain"
	"gorm.io/gorm"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return BrandRepository{db: db}
}

func (br BrandRepository) Create(brand domain.Brand) error {
	return br.db.Create(&brand).Error
}
