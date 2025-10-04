package mac

type MacBatten struct{}
type MacImage struct{}

func (s *MacImage) Render() string {
	return "Mac logic for rendering image"
}

func (s *MacBatten) Render() string {
	return "Mac Logic for rendering Batten"
}
