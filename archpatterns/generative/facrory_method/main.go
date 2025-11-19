package main

import "fmt"

func main() {
	productShoes, _ := getProduct(concreteProductShoes)
	fmt.Println(productShoes.Name(), productShoes.Price())

	productShirt, _ := getProduct(concreteProductShirt)
	fmt.Println(productShirt.Name(), productShirt.Price())
}
