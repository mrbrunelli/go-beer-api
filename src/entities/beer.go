package entities

import "strconv"

// Para deixar público a struct, function, interface devem começar com letras maiúsculas

type Beer struct {
	Name  string
	Price float64
}

func (b Beer) GetPriceInCurrency() string {
	return strconv.FormatFloat(b.Price, 'f', 2, 64)
}
