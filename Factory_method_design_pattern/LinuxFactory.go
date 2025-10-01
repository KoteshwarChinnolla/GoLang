package main

// Linux Factory
type LinuxFactory struct{}

func (sf *LinuxFactory) RenderImage() Image {
	return &LinuxImage{}
}

func (sf *LinuxFactory) RenderButten() Butten {
	return &LinuxButten{}
}
