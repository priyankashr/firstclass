package loop

import "fmt"

func LoopStatement() {

	sum := 0
	i := 0
	for i < 10 {
		sum += i // sum = sum + i
		i++      // i = i + 1
	}
	fmt.Println(sum)
}

func GetPrimeNumbers(last int) { // 2,3, 5 ...
	// var i, j int
	// for i = 2; i < last; i++ {
	// 	for j = 2; j <= (i / j); j++ { // eg. if i=10  and j = 2, j=10/2
	// 		if i%j == 0 {
	// 			break
	// 		}
	// 	}
	// 	if j > (i / j) {
	// 		fmt.Printf("%d is prime\n", i)
	// 	}
	// }
	plus, minus := PlusMinus(1, 2)
	fmt.Println("a + b = ", plus)
	fmt.Println("a - b = ", minus)

}

func PlusMinus(num1 int, num2 int) (int, int) {
	plus := num1 + num2
	minus := num1 - num2
	return plus, minus
}

func CheckSwap() {
	var a, b int
	a = 5
	b = 10
	fmt.Printf("Before swapping, a = %d\n", a)
	fmt.Printf("Before swapping, b = %d\n", b)
	swapByRef(&a, &b) // call by value case (5, 10) // call by reference case (0x3434, 0x234)
	fmt.Printf("After swapping, a = %d\n", a)
	fmt.Printf("After swapping, b = %d\n", b)
}

// Call by Reference
func swapByRef(x *int, y *int) int {
	var temp int
	temp = *x // save x to temporary variable
	*x = *y   // put y into x
	*y = temp // put temp into y
	return temp
}

// Call by value
// func swap(x int, y int) int {
// 	var temp int
// 	temp = x // save x to temporary variable
// 	x = y    // put y into x
// 	y = temp // put temp into y
// 	fmt.Printf("Inside swap function swapping, x = %d\n", x)
// 	fmt.Printf("Inside swap function  swapping, y = %d\n", y)
// 	return temp
