package main

import "fmt"

type concreteOsFactory int

const (
	concreteOsFactoryMac concreteOsFactory = iota
	concreteOsFactoryWindows
)

func getOsFactory(factory concreteOsFactory) (OsFactory, error) {
	switch factory {
	case concreteOsFactoryMac:
		return NewMacOsFactory(), nil
	case concreteOsFactoryWindows:
		return NewWindowsOsFactory(), nil
	default:
		return nil, fmt.Errorf("factory not supported")
	}
}
