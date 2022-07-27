package repositories

import (
	"github.com/mrbrunelli/go-beer-api/src/entities"
)

type BeerRepository interface {
	FindAll() []entities.Beer
	FindById(id string) entities.Beer
	Create(name string, price float64) bool
	Count() int
}
