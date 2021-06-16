package goexamples

import (
	"fmt"
	"time"
)

// "time"

type User struct {
	name  string
	email string
}

//embedding
type Admin struct {
	User
	email    string
	keyrings string
}

func Examples1() {
	admin := Admin{
		User:     User{name: "example1", email: "example@gmail.com"},
		email:    "admin@gmail.com",
		keyrings: "my keys",
	}
	fmt.Println(admin.name)
	fmt.Println(admin.email)
	fmt.Println(admin.email)

}

//channels-- connection between concurrent goroutines
func Examples2() {
	// msg := make(chan string, 3)

	//goroutines
	// go func() {
	// 	msg <- "example"

	// }()

	// go message(msg)

	// message := <-msg
	// fmt.Println(message)
	// fmt.Println(message)
	// fmt.Println(message)

	// done := make(chan bool, 1)
	// go worker(done)
	// <-done

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		time.Sleep(2 * time.Second)
		c1 <- "passing two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "passing one"
	}()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	// default:
	// 	// 	fmt.Println("default")
	// 	case <-time.After(4 * time.Second):
	// 		fmt.Println("timeout")
	// 	}
	// }

	// }

	// func worker(done chan bool) {
	// 	fmt.Println("working.....")
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Println("done")
	// 	done <- true
	// }

	// func message(channel chan string) {
	// 	channel <- "first data"
	// 	channel <- "second data"
	// 	channel <- "third data"

	// closingChannel()
	// rangeOnChannel()

	exampleTimer()
}

func exampleTimer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1, fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}
func rangeOnChannel() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for value := range queue {
		fmt.Println(value)

	}
}

func closingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			rec, close := <-jobs //close boolean
			if close {
				fmt.Println("received jobs= ", rec)
			} else {
				fmt.Println("all jobs finished")
				done <- true
				return

			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job =", i)
	}
	close(jobs)
	fmt.Println("closed")
	<-done
}
