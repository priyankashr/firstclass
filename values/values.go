package values

import "fmt"

//Variable usage and decleration
func Variable() {
	var a, b, c = 5, 6, "djdjd"
	x := 100

	a += 5 // a = a + 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)

}

//Operators usage
func Values() {
	fmt.Println("First" + "Last")
	fmt.Println("10-1 = ", 10-1)
	fmt.Println("10+1 = ", 10+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
