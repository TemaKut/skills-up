package main

type File struct {
	size int
}

func NewFile(size int) *File {
	return &File{size: size}
}

func (f *File) Size() int {
	return f.size
}
