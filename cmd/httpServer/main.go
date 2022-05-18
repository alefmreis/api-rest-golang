package main

import (
	"fmt"

	"github.com/alefmreis/api-rest-golang/infraestructure"
	"github.com/alefmreis/api-rest-golang/src/controllers"
	"github.com/alefmreis/api-rest-golang/src/core/services"
	"github.com/alefmreis/api-rest-golang/src/helpers"
	"github.com/alefmreis/api-rest-golang/src/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := infraestructure.NewEnv()
	db := infraestructure.NewDatabase(env)
	uuidGenerator := helpers.NewUUID()

	brandRepository := repositories.NewBrandRepository(db)
	brandCreateUseCase := services.NewBrandService(brandRepository, uuidGenerator)
	brandController := controllers.NewBrandController(brandCreateUseCase)

	router := gin.New()

	v1 := router.Group("/api/v1")

	v1.POST("/brands", brandController.Create)

	router.Run(fmt.Sprintf(":%s", env.APPLICATION_PORT))

}
