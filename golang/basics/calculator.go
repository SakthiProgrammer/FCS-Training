package basics

import "fmt"

func Calculator(){

	var a, b int
	var o string

	var again = true

	for again {
		fmt.Printf("Enter Number 1: ")
		fmt.Scan(&a)
		fmt.Printf("Enter Number 2: ")
		fmt.Scan(&b)
		fmt.Printf("Enter Operator: ")
		fmt.Scan(&o)
		// fmt.Println(o)
		// ans := calculator(a, b, o)
		// fmt.Println("Ans: ",ans)
		switch o {
			case "+" :
				fmt.Println("Addition..")
				fmt.Println(a + b)
			case "-":
				fmt.Println("Subtraction..")
				fmt.Println(a - b)
			case "*":
				fmt.Println("Multiplication..")
				fmt.Println(a * b)
			case "/":
				fmt.Println("Divition..")
				fmt.Println(a / b)
		}

		var in string
		fmt.Printf("Do You want to Calculate Again: (Y/n) ")
		fmt.Scan(&in)
		if(in == "y" || in == "Y"){
			again = true
		}else{
			again = false
		}

	}

	fmt.Println("Exited..")

}

// func calculator(a int, b int, o rune) int {
	
// }