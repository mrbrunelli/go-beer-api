package services

import (
	"github.com/mrbrunelli/go-beer-api/src/repositories"
)

type NewBeer struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateBeerService struct {
	Repo repositories.BeerRepository
}

func (s CreateBeerService) Execute(b NewBeer) bool {
	return s.Repo.Create(b.Name, b.Price)
}
