package main

import "fmt"

func main() {
	macOsFactory, _ := getOsFactory(concreteOsFactoryMac)

	macOsBrowser := macOsFactory.MakeDefaultBrowser()
	fmt.Println(macOsBrowser.Name())

	macOsTerminal := macOsFactory.MakeTerminal()
	fmt.Println(macOsTerminal.Name(), macOsTerminal.Exec("echo 'Hello World'"))

	windowsOsFactory, _ := getOsFactory(concreteOsFactoryWindows)
	windowsOsBrowser := windowsOsFactory.MakeDefaultBrowser()
	fmt.Println(windowsOsBrowser.Name())

	windowsOsTerminal := windowsOsFactory.MakeTerminal()
	fmt.Println(windowsOsTerminal.Name(), windowsOsTerminal.Exec("echo 'Hello World'"))
}
