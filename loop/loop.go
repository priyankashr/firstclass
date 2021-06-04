package loop

import "fmt"

func LoopStatement() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i // sum = sum + i
		i++      //i = i + 1
	}
	fmt.Println(sum)
}

// try nested for loop-- 2 for loops and then apply break:
// 1. nested for loop
// 2. can break statement exites the both loop at once or not

// 3. if not then how can we break both loop
