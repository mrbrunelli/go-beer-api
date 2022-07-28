package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrbrunelli/go-beer-api/src/services"
)

type BeerController struct {
	CreateBeerService services.CreateBeerService
	FindBeersService  services.FindBeersService
	FindBeerService   services.FindBeerService
}

func (c BeerController) Create(w http.ResponseWriter, r *http.Request) {
	var newBeer services.NewBeer
	_ = json.NewDecoder(r.Body).Decode(&newBeer)

	response := c.CreateBeerService.Execute(newBeer)
	json.NewEncoder(w).Encode(response)
}

func (c BeerController) FindAll(w http.ResponseWriter, r *http.Request) {
	response := c.FindBeersService.Execute()
	json.NewEncoder(w).Encode(response)
}

func (c BeerController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := c.FindBeerService.Execute(params["id"])
	json.NewEncoder(w).Encode(response)
}
