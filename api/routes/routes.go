package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)


type Route struct {
	Path string
	Method string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, productRoutes ProductRoutes, categoryRoutes CategoryRoutes) {
	
	allRoutes := categoryRoutes.Routes()
	allRoutes = append(allRoutes, productRoutes.Routes()...)
	

	for _, route := range allRoutes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}