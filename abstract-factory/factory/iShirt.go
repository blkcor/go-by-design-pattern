package factory

type IShirt interface {
	SetLogo(logo string)
	SetSize(size string)
	GetLogo() string
	GetSize() string
}

type Shirt struct {
	logo string
	size string
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size string) {
	s.size = size
}

func (s *Shirt) GetSize() string {
	return s.size
}
