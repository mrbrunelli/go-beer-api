package entities

import "testing"

// Para executar os testes de todos os diretórios, precisamos executar "go test ./..."
// Coletar cobertura "go test ./... -coverprofile=coverage.out"
// Exibir coverage em html "go tool cover -html=coverage.out"

func TestPriceInCurrency(t *testing.T) {
	got := Beer{
		Id:    "1",
		Name:  "Spaten",
		Price: 4.99,
	}.GetPriceInCurrency()
	want := "4,99"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestManyPriceInCurrency(t *testing.T) {
	type testCase struct {
		arg1     float64
		expected string
	}

	testCases := []testCase{
		{4.99, "4,99"},
		{0.98, "0,98"},
		{2.45, "2,45"},
		{13.47, "13,47"},
		{102.55, "102,55"},
	}

	for _, test := range testCases {
		got := Beer{"1", "Nice beer", test.arg1}.GetPriceInCurrency()
		want := test.expected

		if got != want {
			t.Errorf("Output %q not equal to expected %q", got, want)
		}
	}
}

func TestNameInUppercase(t *testing.T) {
	beer := Beer{"1", "Colorado", 13.50}

	got := beer.GetNameInUpperCase()
	want := "COLORADO"

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func TestToJson(t *testing.T) {
	beer := Beer{"1", "Heineken", 6.19}

	got := beer.ToJson()
	want := `{"id":"1","name":"Heineken","price":6.19}`

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
