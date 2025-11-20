package main

func newLink[T any](v *T) *T {
	if v == nil {
		return nil
	}

	newValue := *v
	
	return &newValue
}
