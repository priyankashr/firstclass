package rangeexample

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
}

func newPerson(firstName string, lastName string) *person {
	p := person{
		firstName: firstName,
		lastName:  lastName,
	}
	p.age = 29
	return &p
}

func RangeExample() {

	person1 := person{
		firstName:   "Priyanka",
		lastName:    "Shrestha",
		phoneNumber: "9999999",
		age:         29}

	if person1.age < 30 {
		fmt.Println("firstname = ", person1.firstName)
	} else {
		fmt.Println("fistname = ", person1.lastName)
	}

	person2 := person{
		firstName: "subrat"}
	fmt.Println("Person = ", person2)

	fmt.Println(newPerson("ram", "gautam"))
	// nums := []int{2, 3, 4, 5, 6}
	// exampleOne(nums)
	// exampleMap()
	// exampleTwo()
	// exampleVariadicFunction()
}

func exampleVariadicFunction() {
	nums := []int{2, 3, 4, 5, 6}
	sum(1, 2)
	sum(1, 2, 6, 7)
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println("nums value = ", nums)
	sum := 0
	for _, num := range nums { // num = 2, 3, 4, 5
		sum += num
	}
	fmt.Println("sum = ", sum)
}

func exampleTwo() {
	for i, value := range "ABCabc" {
		fmt.Println(i, value)
	}
}

func exampleMap() {
	// map  key => value  name => "subrat"
	// map key is always unique

	// n := map[string]int{"one": 1, "two": 2}
	// fmt.Println("map = ", n)

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	delete(m, "three")
	fmt.Println("map = ", m)

	for key, value := range m {
		fmt.Printf("%v = %v\n", key, value)
	}
}

func exampleOne(nums []int) {
	// nums := []int{2, 3, 4, 5, 6}
	sum := 0
	// sum = nums[0] + nums[1] + nums[2] + nums[3]
	for i := 0; i < len(nums); i++ {
		sum += nums[i] // sum = sum + nums[i]
	}
	fmt.Println("sum = ", sum)

	sum = 0
	//  range => (index,value)
	// _ => blank identifier
	for _, num := range nums { // num = 2, 3, 4, 5
		sum += num
	}
	fmt.Println("sum = ", sum)
}
