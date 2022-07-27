package repositories

import "github.com/mrbrunelli/go-beer-api/src/entities"

type BeerRepositoryInMemory struct {
	beers []entities.Beer
}

func (r BeerRepositoryInMemory) FindAll() []entities.Beer {
	return r.beers
}

func (r *BeerRepositoryInMemory) Create(name string, price float64) bool {
	beer := entities.Beer{Name: name, Price: price}

	oldLen := r.Count()
	r.beers = append(r.beers, beer)
	newLen := r.Count()

	return newLen > oldLen
}

func (r BeerRepositoryInMemory) Count() int {
	return len(r.beers)
}
