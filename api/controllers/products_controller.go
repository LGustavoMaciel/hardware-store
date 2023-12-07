package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"hardware-store/api/models"
	"hardware-store/api/repository"
	"hardware-store/api/utils"
)


type ProductsController interface {
	PostProduct(http.ResponseWriter, *http.Request)
} 

type productsControllerImpl struct {
	productsRepository repository.ProductsRepository
}

func NewProductsController(productsRepository repository.ProductsRepository) *productsControllerImpl {
	return &productsControllerImpl{productsRepository}
}



func (c *productsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	// product.CheckStatus()

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}