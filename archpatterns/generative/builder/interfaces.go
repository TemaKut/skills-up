package main

type StringBuilder interface {
	SetStar()
	SetNumber(uint65 uint64)
	SetString(str string)
	Build() string
}
