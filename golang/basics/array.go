package main

import (
	"log"
	"os"
	"time"
)

func ll() {
	log.Println("hello world")
}

func main() {

	log.Println("Server Started")
	f, err := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	ll()

	log.Println("M001-(+)")
	var arr [2]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 3
	}
	log.Println(arr)

	var arr1 = [3]int{12, 32, 12}
	log.Println(arr1)

	arr2 := []int{12, 213, 123, 1, 27, 28, 34, 17, 90}
	log.Println(arr2)

	slice := arr2[2:6]
	log.Println(slice)

	index := 2
	slice1 := slice[:index]
	slice2 := slice[index+1:]
	slice1 = append(slice1, slice2...)
	log.Println(slice1)

	decSlice := make([]int, 2)
	log.Println(decSlice)
	log.Println("M001-(-)")

}
