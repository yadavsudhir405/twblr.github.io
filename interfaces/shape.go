package interfaces

type Hybrid []Shape

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

func (h Hybrid) Area() int{
	area:=0
	for _,shape := range h{
		area+=shape.Area()
	}
	return area
}

type Hybrids struct{
	shapes []Shape
}
func (h *Hybrids) Area()int{
	area:=0
	for _,shape := range h.shapes{
		area+=shape.Area()
	}
	return area
}