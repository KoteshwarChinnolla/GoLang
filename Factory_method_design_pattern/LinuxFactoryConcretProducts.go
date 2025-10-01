package main

type LinuxButten struct{}
type LinuxImage struct{}

func (s *LinuxImage) Render() string {
	return "Rendering a linux image"
}

func (s *LinuxButten) Render() string {
	return "Rendering a linux butten"
}
