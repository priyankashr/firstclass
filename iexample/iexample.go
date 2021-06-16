package iexample

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type geometry interface {
	area() float64
	perim() (float64, error)
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 { //pir^2
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() (float64, error) { //2Pir
	return 2 * math.Pi * c.radius, nil
}

func (r rect) perim() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0, errors.New("width cannot be 0")
	}
	return 2*r.width + 2*r.height, nil
}

// using interface
func measure(g geometry) (float64, float64, error) {
	perm, err := g.perim()
	if err != nil {
		return 0, 0, err
	}
	return g.area(), perm, nil

}

func MainIStart() {
	r := rect{width: 0, height: 5}
	// c := circle{radius: 5}
	// fmt.Println("rectange area = ", r.area())
	// fmt.Println("rectange perimeter = ", r.perim())
	// fmt.Println("circle area = ", c.area())
	// fmt.Println("circle perimeter = ", c.perim())
	a, b, err := measure(r)
	if err != nil {
		fmt.Println("%w", err)
		// log.Fatal(err)
	}
	fmt.Println(a)
	fmt.Println(b)
	// measure(c)
}

func GoRoutingExample() {
	Routine("direct")

	go Routine("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going.....")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func Routine(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
