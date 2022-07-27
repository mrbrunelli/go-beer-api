package entities

import "testing"

// Para executar os testes de todos os diretórios, precisamos executar "go test ./..."

func TestPriceInCurrency(t *testing.T) {
	got := Beer{
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
		got := Beer{"Nice beer", test.arg1}.GetPriceInCurrency()
		want := test.expected

		if got != want {
			t.Errorf("Output %q not equal to expected %q", got, want)
		}
	}
}

func TestNameInUppercase(t *testing.T) {
	beer := Beer{"Colorado", 13.50}

	got := beer.GetNameInUpperCase()
	want := "COLORADO"

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
