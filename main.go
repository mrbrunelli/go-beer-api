package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrbrunelli/go-beer-api/src/controllers"
	"github.com/mrbrunelli/go-beer-api/src/repositories"
	"github.com/mrbrunelli/go-beer-api/src/services"
)

func main() {
	beerRepository := repositories.BeerRepositoryInMemory{}
	createBeerService := services.CreateBeerService{Repo: &beerRepository}
	beerController := controllers.BeerController{CreateBeerService: createBeerService}

	router := mux.NewRouter()
	router.HandleFunc("/beer", beerController.Create).Methods("POST")

	log.Println("Start server")
	log.Fatal(http.ListenAndServe(":3000", router))
}
