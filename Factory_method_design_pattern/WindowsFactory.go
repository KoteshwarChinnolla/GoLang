package main

// Windows Factory
type WindowsFactory struct{}

func (sf *WindowsFactory) RenderImage() Image {
	return &WindowsImage{}
}

func (sf *WindowsFactory) RenderButten() Butten {
	return &WindowsButten{}
}
