package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"hardware-store/api/controllers"
	"hardware-store/api/database"
	"hardware-store/api/repository"
	"hardware-store/api/routes"
)

var (
	port = flag.Int("p", 5000, "Set port")
)

func Run () {
	flag.Parse()
	db := database.Connect()
	if db == nil {
		defer db.Close()
	}

	categoriesRepository := repository.NewCategoriesRepository(db)
	productsRepository := repository.NewProductsRepository(db)


	categoriesController := controllers.NewCategoriesRepository(categoriesRepository)
	productsController := controllers.NewProductsController(productsRepository)


	categoryRoutes := routes.NewCategoryRoutes(categoriesController)
	productsRoutes := routes.NewProductRoutes(productsController)

	router := mux.NewRouter().StrictSlash(true)

	routes.Install(router, categoryRoutes, productsRoutes)


	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	

	fmt.Println("Api Listening", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))

}