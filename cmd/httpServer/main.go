package main

import (
	"fmt"

	"github.com/alefmreis/api-rest-golang/infraestructure"
	"github.com/alefmreis/api-rest-golang/src/controllers"
	ports "github.com/alefmreis/api-rest-golang/src/core/ports/repositories"
	"github.com/alefmreis/api-rest-golang/src/core/services"
	"github.com/alefmreis/api-rest-golang/src/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	env := infraestructure.NewEnv()
	db := infraestructure.NewDatabase(env)
	// uuidGenerator := helpers.NewUUID()

	brandController := initBrand(db)

	// brandRepository := repositories.NewBrandRepository(db)
	// brandCreateUseCase := services.NewBrandService(brandRepository, uuidGenerator)
	// brandController := controllers.NewBrandController(brandCreateUseCase)

	router := gin.New()

	v1 := router.Group("/api/v1")

	v1.POST("/brands", brandController.Create)

	router.Run(fmt.Sprintf(":%s", env.APPLICATION_PORT))

}

func initBrand(db *gorm.DB) controllers.BrandController {
	wire.Build(
		wire.Bind(new(ports.BrandRepository), new(*repositories.BrandRepository)),
		services.NewBrandService,
		controllers.NewBrandController)

	return controllers.BrandController{}
}
