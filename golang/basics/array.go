package main

import "fmt"

func main() {
	var arr [2]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 3
	}
	fmt.Println(arr)

	var arr1 = [3]int{12, 32, 12}
	fmt.Println(arr1)

	arr2 := []int{12, 213, 123, 1, 27, 28, 34, 17, 90}
	fmt.Println(arr2)

	slice := arr2[2:6]
	fmt.Println(slice)

	index := 2
	slice1 := slice[:index]
	slice2 := slice[index+1:]
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	decSlice := make([]int, 2)
	fmt.Println(decSlice)
}
