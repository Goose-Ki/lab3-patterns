package abstract_factory

import "fmt"

type Button interface {
	Paint() string
}

type Checkbox interface {
	Paint() string
}
type WindowsButton struct{}

func (wb *WindowsButton) Paint() string {
	return "Windows style button"
}

type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Paint() string {
	return "Windows style checkbox"
}

type MacOSButton struct{}

func (mb *MacOSButton) Paint() string {
	return "MacOS style button"
}

type MacOSCheckbox struct{}

func (mc *MacOSCheckbox) Paint() string {
	return "MacOS style checkbox"
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (wf *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type MacOSFactory struct{}

func (mf *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mf *MacOSFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}

func CreateUI(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	fmt.Println(button.Paint())
	fmt.Println(checkbox.Paint())
}
