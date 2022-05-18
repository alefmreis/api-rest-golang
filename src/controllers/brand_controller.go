package controllers

import (
	"net/http"

	"github.com/alefmreis/api-rest-golang/src/core/domain"
	services "github.com/alefmreis/api-rest-golang/src/core/ports/services"
	"github.com/gin-gonic/gin"
)

type BrandController struct {
	Service services.BrandService
}

func NewBrandController(brandService services.BrandService) BrandController {
	return BrandController{Service: brandService}
}

func (bc BrandController) Create(c *gin.Context) {
	var brand domain.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	bc.Service.Save(brand.Name)

	c.JSON(201, nil)
}
