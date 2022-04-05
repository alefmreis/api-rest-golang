package ports

type BrandService interface {
	Save(name string) error
}
