package main

func main() {
	var realJSONStorage IJSONDataStorage = NewJSONDataStorage()
	println(realJSONStorage.FetchJSONData())

	var adapterXMLToJSONStorage IJSONDataStorage = NewXMLToJSONStorageAdapter(NewXMLDataStorage())
	println(adapterXMLToJSONStorage.FetchJSONData())
}
