package array

import (
	"fmt"
	"reflect"
)

func ArrayMain() {
	simpleArray()
	twoDimensionalArray()
}

func twoDimensionalArray() {
	td := [2][3]int{{1, 2, 5}, {3, 4}}
	fmt.Println(td)

	var twoDArr [2][3]int
	for i := 0; i < len(twoDArr); i++ { // i -- 0, 1
		for j := 0; j < len(twoDArr[i]); j++ { // j -- 0, 1, 2
			twoDArr[i][j] = i + j // twoDArr[0][0] , [0][1],[0][2] .. [1][0]. [1][1]
		}
	}

	fmt.Println("twoDarr value = ", twoDArr)

}

func simpleArray() {
	var arr [5]int //[0,1,2,3,4]
	fmt.Printf("Value of arr %v\n", arr)
	fmt.Printf("Type of arr %T\n", arr)
	fmt.Println(reflect.TypeOf(arr).Kind())

	arr[4] = 4
	fmt.Println("setting the value", arr[4])
	fmt.Println("len of arr = ", len(arr))

	arr2 := [5]int{4, 5, 2, 6, 8}
	fmt.Println("value of arr2", arr2)
}
