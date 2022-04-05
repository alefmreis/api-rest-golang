package ports

import "github.com/alefmreis/api-rest-golang/src/core/domain"

type BrandRepository interface {
	Create(domain.Brand) error
}
