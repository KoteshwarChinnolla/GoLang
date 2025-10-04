package Mac

type MacBatten struct{}
type MacImage struct{}

func (s *MacImage) Render() string {
	return "Rendering a Mac image"
}

func (s *MacBatten) Render() string {
	return "Rendering a Mac Batten"
}
