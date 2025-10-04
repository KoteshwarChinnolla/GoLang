package windows

// Concrete Products
type WindowsButten struct{}
type WindowsImage struct{}

func (s *WindowsImage) Render() string {
	return "Rendering a windows image"
}

func (s *WindowsButten) Render() string {
	return "Rendering a windows butten"
}
