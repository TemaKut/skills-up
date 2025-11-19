package main

type Shirt struct{}

func NewShirt() *Shirt {
	return &Shirt{}
}

func (s Shirt) Name() string {
	return "Shirt"
}

func (s Shirt) Price() uint64 {
	return 721
}

type Shoes struct{}

func NewShoes() *Shoes {
	return &Shoes{}
}

func (s Shoes) Name() string {
	return "Shoes"
}

func (s Shoes) Price() uint64 {
	return 1213
}
