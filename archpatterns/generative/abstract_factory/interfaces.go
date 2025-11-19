package main

type OsFactory interface {
	MakeTerminal() Terminal
	MakeDefaultBrowser() Browser
}

type Terminal interface {
	Name() string
	Exec(cmd string) error
}

type Browser interface {
	Name() string
}
