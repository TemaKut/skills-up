package main

type BackendDeveloper struct {
}

func NewBackendDeveloper() *BackendDeveloper {
	return &BackendDeveloper{}
}

func (b *BackendDeveloper) Skills() []string {
	return []string{
		"Golang",
		"Docker",
		"k8s",
		"Nats",
		"Kafka",
		"Postgres",
	}
}
