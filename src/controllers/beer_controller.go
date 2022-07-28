package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mrbrunelli/go-beer-api/src/services"
)

type BeerController struct {
	CreateBeerService services.CreateBeerService
}

func (c BeerController) Create(w http.ResponseWriter, r *http.Request) {
	var newBeer services.NewBeer
	_ = json.NewDecoder(r.Body).Decode(&newBeer)

	response := c.CreateBeerService.Execute(newBeer)
	json.NewEncoder(w).Encode(response)
}
