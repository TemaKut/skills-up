package main

func main() {
	rootFolder := NewFolder(
		NewFile(100),
		NewFile(200),
		NewFolder(
			NewFile(100),
			NewFolder(
				NewFile(100),
			),
		),
	)

	println(rootFolder.Size())
}
