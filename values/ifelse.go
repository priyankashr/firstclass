package values

import "fmt"

func IfElseGo() {

	// var x int
	// fmt.Println("Enter the x =")
	// fmt.Println(&x)

	if x := 10; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println(fmt.Println("x is odd"))
	}
}

func Swithcase() {
	x := 1
	switch {
	case x > 100:
		fmt.Println("case > 100")
	case x > 10:
		fmt.Println("case > 10")
	default:
		fmt.Println("default")

	}
}
