package main

type WindowsTerminal struct {
}

func NewWindowsTerminal() *WindowsTerminal {
	return &WindowsTerminal{}
}

func (w WindowsTerminal) Name() string {
	return "Windows Terminal"
}

func (w WindowsTerminal) Exec(cmd string) error {
	// do something..
	return nil
}

type MacOsTerminal struct {
}

func NewMacOsTerminal() *MacOsTerminal {
	return &MacOsTerminal{}
}

func (m MacOsTerminal) Name() string {
	return "Mac OS Terminal"
}

func (m MacOsTerminal) Exec(cmd string) error {
	// do something..
	return nil
}
