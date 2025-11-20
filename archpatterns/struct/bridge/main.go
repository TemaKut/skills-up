package main

func main() {
	redApple := NewAppleProduct(NewRedColor())
	println(redApple.Name())

	greenApple := NewAppleProduct(NewGreenColor())
	println(greenApple.Name())
}
