package main

import "fmt"

type AppleProduct struct {
	color Color
}

func NewAppleProduct(color Color) *AppleProduct {
	return &AppleProduct{
		color: color,
	}
}

func (a AppleProduct) Name() string {
	return fmt.Sprintf("i`m an apple with %s color", a.color.Name())
}
