package main

type MacOsBrowser struct {
}

func NewMacOsBrowser() *MacOsBrowser {
	return &MacOsBrowser{}
}

func (m MacOsBrowser) Name() string {
	return "MacOsBrowser"
}

type WindowsBrowser struct {
}

func NewWindowsBrowser() *WindowsBrowser {
	return &WindowsBrowser{}
}

func (w WindowsBrowser) Name() string {
	return "WindowsBrowser"
}
