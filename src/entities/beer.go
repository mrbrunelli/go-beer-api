package entities

import (
	"strconv"
	"strings"
)

// Para deixar público a struct, function, interface devem começar com letras maiúsculas

type Beer struct {
	Name  string
	Price float64
}

func (b Beer) GetPriceInCurrency() string {
	return strings.Replace(strconv.FormatFloat(b.Price, 'f', 2, 64), ".", ",", 1)
}

func (b Beer) GetNameInUpperCase() string {
	return strings.ToUpper(b.Name)
}
