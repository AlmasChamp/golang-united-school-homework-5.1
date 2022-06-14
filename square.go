package square

type Square struct {
	i int
}

func (s *Square) Area() int {
	return s.i * s.i
}

func (s *Square) Perimetr() int {
	return s.i * 4
}
