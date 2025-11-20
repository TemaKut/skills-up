package main

import "fmt"

func main() {
	message := NewMessage("Hello World")
	messageClone := message.Clone()

	message.Text = "Updated text"

	fmt.Println(message)
	fmt.Println(messageClone)
}
