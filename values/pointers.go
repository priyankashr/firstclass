package values

import "fmt"

func PointersValue() {
	var p *int
	i := 20

	p = &i
	fmt.Println(*p)
}
