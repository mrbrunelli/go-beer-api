package repositories

import (
	"os"
	"reflect"
	"testing"
)

var repo BeerRepository

func seed(r BeerRepository) {
	r.Create("Boa", 2.49)
	r.Create("Duplo Malte", 3.78)
	r.Create("Colorado", 13.47)
}

// TestMain executa antes de todos os cenários. Parecido com o BeforeAll do Jest
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

func TestFindById(t *testing.T) {
	got := repo.FindById("1")
	isEmpty := reflect.ValueOf(got).IsZero()

	if isEmpty {
		t.Errorf("Beer not found. Got %v, isEmpty %t", got, isEmpty)
	}

}

// Testa o desempenho de um caso. Execute "go test -bench=." ou pelo nome "go test -bench=Create"
func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repo.Create("Nice beer", 12.44)
	}
}

func BenchmarkFindById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repo.FindById("200")
	}
}
