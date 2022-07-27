package repositories

import (
	"os"
	"testing"
)

var repo BeerRepository

func seed(r BeerRepository) {
	r.Create("Boa", 2.49)
	r.Create("Duplo Malte", 3.78)
	r.Create("Colorado", 13.47)
}

func TestMain(m *testing.M) {
	repo = &BeerRepositoryInMemory{}
	seed(repo)
	code := m.Run()
	os.Exit(code)
}

func TestCreateBeer(t *testing.T) {
	got := repo.Create("Spaten", 6.20)
	want := true

	if got != want {
		t.Errorf("Got %t want %t", got, want)
	}
}

func TestFindAllBeers(t *testing.T) {
	got := repo.FindAll()
	want := 4

	if len(got) != want {
		t.Errorf("Got %v, want %v", len(got), want)
	}
}
