package interfacesandstructs

type Rectangle struct {
	Height float64
	Width  float64
}

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return height * width
}
