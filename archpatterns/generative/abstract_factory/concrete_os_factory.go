package main

type MacOsFactory struct {
}

func NewMacOsFactory() *MacOsFactory {
	return &MacOsFactory{}
}

func (m MacOsFactory) MakeTerminal() Terminal {
	return NewMacOsTerminal()
}

func (m MacOsFactory) MakeDefaultBrowser() Browser {
	return NewMacOsBrowser()
}

type WindowsOsFactory struct{}

func NewWindowsOsFactory() *WindowsOsFactory {
	return &WindowsOsFactory{}
}

func (w WindowsOsFactory) MakeTerminal() Terminal {
	return NewWindowsTerminal()
}

func (w WindowsOsFactory) MakeDefaultBrowser() Browser {
	return NewWindowsBrowser()
}
