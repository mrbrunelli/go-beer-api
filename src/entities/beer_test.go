package entities

import "testing"

// Para executar os testes de todos os diretórios, precisamos executar "go test ./..."

func TestPriceInCurrency(t *testing.T) {
	got := Beer{
		Name:  "Spaten",
		Price: 4.99,
	}.GetPriceInCurrency()
	want := "4.99"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
