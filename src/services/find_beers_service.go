package services

import (
	"github.com/mrbrunelli/go-beer-api/src/entities"
	"github.com/mrbrunelli/go-beer-api/src/repositories"
)

type FindBeersService struct {
	Repo repositories.BeerRepository
}

func (s FindBeersService) Execute() []entities.Beer {
	return s.Repo.FindAll()
}
