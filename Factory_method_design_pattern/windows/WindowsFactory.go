package windows

import (
	"Factory_method_design_pattern/interfaces"
)

// Windows Factory
type WindowsFactory struct{}

func (sf *WindowsFactory) RenderImage() interfaces.Image {
	return &WindowsImage{}
}

func (sf *WindowsFactory) RenderBatten() interfaces.Batten {
	return &WindowsBatten{}
}
