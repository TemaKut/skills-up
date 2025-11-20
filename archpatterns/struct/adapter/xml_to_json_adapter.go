package main

import "fmt"

type XMLToJSONStorageAdapter struct {
	xmlStorage IXMLDataStorage
}

func NewXMLToJSONStorageAdapter(xmlStorage IXMLDataStorage) IJSONDataStorage {
	return &XMLToJSONStorageAdapter{
		xmlStorage: xmlStorage,
	}
}

// FetchJSONData Just xml string wrap to json
func (x *XMLToJSONStorageAdapter) FetchJSONData() string {
	return fmt.Sprintf("[\"%s\"]", x.xmlStorage.FetchXMLData())
}
