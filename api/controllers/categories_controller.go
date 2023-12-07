package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"hardware-store/api/models"
	"hardware-store/api/repository"
	"hardware-store/api/utils"
)


type CategoriesController interface {
	PostCategory(http.ResponseWriter, *http.Request)
}

type categoriesControllerImpl struct {
	categoriesRepository repository.CategoriesRepository
}

func NewCategoriesRepository(categoriesRepository repository.CategoriesRepository) *categoriesControllerImpl {
	return &categoriesControllerImpl{categoriesRepository}
}

func (c *categoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err = c.categoriesRepository.Save(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, category.ID))
	utils.WriteAsJson(w, category)

}