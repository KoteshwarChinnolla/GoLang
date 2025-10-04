package mac

import (
	"Factory_method_design_pattern/interfaces"
)

// Mac Factory
type MacFactory struct{}

func (sf *MacFactory) RenderImage() interfaces.Image {
	return &MacImage{}
}

func (sf *MacFactory) RenderBatten() interfaces.Batten {
	return &MacBatten{}
}
