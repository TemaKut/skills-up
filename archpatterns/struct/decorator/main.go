package main

import "fmt"

func main() {
	developerDecorated := NewSQLSkillDecorator(NewBackendDeveloper())

	fmt.Printf("%+v\n", developerDecorated.Skills())
}
