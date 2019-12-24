package interfacesandstructs

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0.0
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
