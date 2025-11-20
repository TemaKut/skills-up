package main

type Folder struct {
	components []Component
}

func NewFolder(components ...Component) *Folder {
	return &Folder{
		components: components,
	}
}

func (f *Folder) Size() int {
	var size int

	for _, component := range f.components {
		size += component.Size()
	}

	return size
}
