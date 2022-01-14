package model

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) GetLogo() string {
	return s.getLogo()
}

func (s *shirt) GetSize() int {
	return s.getSize()
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getSize() int {
	return s.size
}
