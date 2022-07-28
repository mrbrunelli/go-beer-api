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
	findBeersService := services.FindBeersService{Repo: &beerRepository}
	findBeerService := services.FindBeerService{Repo: &beerRepository}

	beerController := controllers.BeerController{
		CreateBeerService: createBeerService,
		FindBeersService:  findBeersService,
		FindBeerService:   findBeerService,
	}

	router := mux.NewRouter()
	router.HandleFunc("/beer", beerController.Create).Methods("POST")
	router.HandleFunc("/beer", beerController.FindAll).Methods("GET")
	router.HandleFunc("/beer/{id}", beerController.FindById).Methods("GET")

	log.Println("Start server")
	log.Fatal(http.ListenAndServe(":3000", router))
}
