package windows

// Concrete Products
type WindowsBatten struct{}
type WindowsImage struct{}

func (s *WindowsImage) Render() string {
	return "Windows logic for rendering images"
}

func (s *WindowsBatten) Render() string {
	return "Windows logic for rendering Batten"
}
