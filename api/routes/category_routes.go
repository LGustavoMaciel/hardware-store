package routes

import (
	"net/http"

	"hardware-store/api/controllers"
)

type CategoryRoutes interface {
	Routes()[]*Route
}

type categoryRoutesImpl struct {
	categoriesController controllers.CategoriesController
}

func NewCategoryRoutes(categoriesController controllers.CategoriesController) *categoryRoutesImpl {
	return &categoryRoutesImpl{categoriesController}
}

func (r *categoryRoutesImpl) Routes() []*Route {
	return []*Route{
		&Route{
			Path:    "/categories",
			Method:  http.MethodPost,
			Handler: r.categoriesController.PostCategory,
		},
	}
}