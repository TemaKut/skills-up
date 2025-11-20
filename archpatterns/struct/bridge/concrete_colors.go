package main

type RedColor struct{}

func NewRedColor() *RedColor {
	return &RedColor{}
}

func (r *RedColor) Name() string {
	return "red"
}

type GreenColor struct{}

func NewGreenColor() *GreenColor {
	return &GreenColor{}
}

func (g GreenColor) Name() string {
	return "green"
}
