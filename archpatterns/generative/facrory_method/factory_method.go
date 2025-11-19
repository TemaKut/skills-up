package main

import "errors"

type concreteProduct int

const (
	concreteProductShirt concreteProduct = iota
	concreteProductShoes
)

func getProduct(concreteProduct concreteProduct) (Product, error) {
	switch concreteProduct {
	case concreteProductShirt:
		return NewShirt(), nil
	case concreteProductShoes:
		return NewShoes(), nil
	default:
		return nil, errors.New("unknown product")
	}
}
