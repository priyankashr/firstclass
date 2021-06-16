package array

import (
	"fmt"
)

func SlicesMain() {
	intro()
	appendValue()
}

func appendValue() {
	s := make([]string, 2)
	fmt.Println("slices = ", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))
	s[0] = "g"
	s[1] = "o"
	fmt.Println("slices = ", s)
	printSlic(s)

	s = append(s, "l")
	fmt.Println("slices = ", s)
	printSlic(s)
}

func intro() {
	// var arr = []int{1, 3, 5}
	// fmt.Printf("Value of arr %v\n", arr)
	// fmt.Printf("Type of arr %T\n", arr)
	// fmt.Printf("length of arr %v\n", len(arr))
	// fmt.Println(reflect.TypeOf(arr).Kind())

	s := make([]string, 6) // make(type, len, cap)
	fmt.Println("slices:", s)

	s[0] = "g"
	s[1] = "o"
	s[2] = "l"
	s[3] = "a"
	s[4] = "n"
	s[5] = "g"
	fmt.Println("slices:", s)
	printSlic(s)

	s = s[:]
	fmt.Println(s)
	printSlic(s)

	s = s[2:4]
	fmt.Println(s)
	printSlic(s)

	s = s[:4]
	fmt.Println(s)
	printSlic(s)

	s = s[2:]
	fmt.Println(s)
	printSlic(s)

	s = s[:]
	fmt.Println(s)

}

func printSlic(s []string) {
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))
}
