package interfaces

type Shape interface {
	Area() int
}
type Square struct{
	side int
}
func (s *Square) Area()int{
	return s.side*s.side
}
type Rectangle struct{
	length int
	breadth int
}
func (r *Rectangle) Area() int{
	return r.length*r.breadth
}
type Hybrid struct{
	square Square
	rectangle Rectangle
}
func (h *Hybrid) Area() int{
	return h.square.Area() +h.rectangle.Area()
}