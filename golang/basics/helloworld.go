package basics

import "fmt"

func Hello(){
	fmt.Printf("Enter Your Name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello ",name)
}