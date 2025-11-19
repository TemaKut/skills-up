package main

import "fmt"

func main() {
	builder := NewRandomPositionStringBuilder()

	builder.SetNumber(1)

	if true {
		builder.SetStar()
	}

	builder.SetString("Hello")

	fmt.Println(builder.Build())
}
