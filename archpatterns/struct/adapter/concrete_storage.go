package main

type JSONDataStorage struct{}

func NewJSONDataStorage() *JSONDataStorage {
	return &JSONDataStorage{}
}

func (J *JSONDataStorage) FetchJSONData() string {
	return `["Hello World!"]`
}

type XMLDataStorage struct{}

func NewXMLDataStorage() *XMLDataStorage {
	return &XMLDataStorage{}
}

func (X *XMLDataStorage) FetchXMLData() string {
	return `<String>Hello world!</String>`
}
