package main

type IJSONDataStorage interface {
	FetchJSONData() string
}

type IXMLDataStorage interface {
	FetchXMLData() string
}
