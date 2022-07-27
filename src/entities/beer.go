package entities

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Para deixar público a struct, function, interface devem começar com letras maiúsculas

type Beer struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (b Beer) GetPriceInCurrency() string {
	return strings.Replace(strconv.FormatFloat(b.Price, 'f', 2, 64), ".", ",", 1)
}

func (b Beer) GetNameInUpperCase() string {
	return strings.ToUpper(b.Name)
}

func (b Beer) ToJson() string {
	j, _ := json.Marshal(b)
	return string(j)
}
