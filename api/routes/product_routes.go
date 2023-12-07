package routes

import (
	"net/http"

	"hardware-store/api/controllers"
)



type ProductRoutes interface {
	Routes() []*Route
}


type productRoutesImpl struct {
	productsController controllers.ProductsController
}


func NewProductRoutes(productsController controllers.ProductsController) *productRoutesImpl {
	return &productRoutesImpl{productsController}
}

func (r *productRoutesImpl) Routes() []*Route {
	return []*Route{
		&Route{
			Path:    "/products",
			Method:  http.MethodPost,
			Handler: r.productsController.PostProduct,
		},
	}
}