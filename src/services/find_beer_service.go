package services

import (
	"github.com/mrbrunelli/go-beer-api/src/entities"
	"github.com/mrbrunelli/go-beer-api/src/repositories"
)

type FindBeerService struct {
	Repo repositories.BeerRepository
}

func (s FindBeerService) Execute(id string) entities.Beer {
	return s.Repo.FindById(id)
}
