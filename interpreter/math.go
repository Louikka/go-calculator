package interpreter

import "github.com/shopspring/decimal"

func Add(v ...float64) float64 {
	acc := decimal.NewFromFloat(0)

	for _, n := range v {
		acc = acc.Add(decimal.NewFromFloat(n))
	}

	r, _ := acc.Float64()
	return r
}

func Sub(v ...float64) float64 {
	argc := len(v)

	if argc == 0 {
		return 0
	}

	if argc == 1 {
		return v[0]
	}

	acc := decimal.NewFromFloat(v[0])

	for _, n := range v[1:] {
		acc = acc.Sub(decimal.NewFromFloat(n))
	}

	r, _ := acc.Float64()
	return r
}

func Mul(v ...float64) float64 {
	argc := len(v)

	if argc == 0 {
		return 0
	}

	if argc == 1 {
		return v[0]
	}

	acc := decimal.NewFromFloat(v[0])

	for _, n := range v[1:] {
		acc = acc.Mul(decimal.NewFromFloat(n))
	}

	r, _ := acc.Float64()
	return r
}
