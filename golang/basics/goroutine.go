// Program to illustrate multiple goroutines

package main

import (
	"fmt"
)

func display(message string) {
	fmt.Println(message)

}

func main() {

	// run two different goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}

	// to sleep main goroutine for 1 sec
	// time.Sleep(time.Second * 1)
}
