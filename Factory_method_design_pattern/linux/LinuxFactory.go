package linux

import (
	"Factory_method_design_pattern/interfaces"
)

// Linux Factory
type LinuxFactory struct{}

func (sf *LinuxFactory) RenderImage() interfaces.Image {
	return &LinuxImage{}
}

func (sf *LinuxFactory) RenderButten() interfaces.Butten {
	return &LinuxButten{}
}
