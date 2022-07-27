package repositories

import (
	"strconv"

	"github.com/mrbrunelli/go-beer-api/src/entities"
)

type BeerRepositoryInMemory struct {
	beers []entities.Beer
}

func (r BeerRepositoryInMemory) FindAll() []entities.Beer {
	return r.beers
}

func (r BeerRepositoryInMemory) FindById(id string) entities.Beer {
	m := map[string]entities.Beer{}

	for _, v := range r.beers {
		m[v.Id] = v
	}

	return m[id]

}

func (r *BeerRepositoryInMemory) Create(name string, price float64) bool {
	oldLen := r.Count()
	id := strconv.Itoa(oldLen + 1)

	beer := entities.Beer{Id: id, Name: name, Price: price}

	r.beers = append(r.beers, beer)
	newLen := r.Count()

	return newLen > oldLen
}

func (r BeerRepositoryInMemory) Count() int {
	return len(r.beers)
}
