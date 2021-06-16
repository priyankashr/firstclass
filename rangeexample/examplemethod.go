package rangeexample

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func MethodExample() {
	rectangle := rect{width: 5, height: 2}
	fmt.Println("Area", rectangle.area())
	fmt.Println("Perimeter", rectangle.perim())

	r := &rectangle
	fmt.Println("Area", r.area())
	fmt.Println("Perimeter", r.perim())
}
