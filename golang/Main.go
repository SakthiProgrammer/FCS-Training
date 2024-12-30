package main

import (
	"encoding/json"
	"fmt"
	stu "learninggo/stuctpack"
)

func main() {
	// ans := B.Palindrome("malayala")
	// fmt.Println(ans)

	stu.Add(1, "Sakthi", "Computer Science", 19, true)
	stu.Add(2, "Ram", " Science", 20, false)
	stu.Add(3, "Thiru", "Mechanical", 21, false)
	stu.Add(4, "Siva", "Civil", 19, true)

	result := stu.Get()
	fmt.Printf("%v", result)
	s, err := json.Marshal(stu.Get())
	if err != nil {
		fmt.Println("Error Found", err)
		return
	}
	fmt.Println("Success to Fetch")
	fmt.Printf("%s", s)
}
