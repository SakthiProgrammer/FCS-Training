package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/*=================1. to create a file ================================ */

	f, err := os.Create("../msg.txt") //../msg.txt - targets root folder inside create msg.txt file

	if err != nil {
		log.Println("Error in createing file")
		return
	}
	defer f.Close()
	fmt.Println(f)

	/*=================2. write data in file =============================== */
	_, err = f.WriteString("Hello world")

	if err != nil {
		log.Println("Error in writing file")
		return
	}

	log.Println("Data writed successfully")

	/*================3. append a data from existing file =================== */

	f, err = os.OpenFile("../msg.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(" append Again data"); err != nil {
		panic(err)
	}

	/*=================4. read a data from existing file===================== */
	f, err = os.Open("../msg.txt")

	if err != nil {
		log.Println("Error in Reading file")
		return
	}

	defer f.Close()

	buffer := make([]byte, 150)

	byteRead, err := f.Read(buffer)

	if err != nil {
		log.Println("Error in Convert byte file")
		return
	}

	fmt.Println(string(buffer[:byteRead]))

	/*
		io/ioutil Package
		=================
	*/

	/*===========1. Write a data in file ==========*/
	data := "Hello Sakthivel"

	err = ioutil.WriteFile("msg.txt", []byte(data), 0655)

	if err != nil {
		log.Println("Error in write file")
		return
	}

	log.Println("Successfully data writed")

	/*===========2. Read a data in file ==========*/

	content, err := ioutil.ReadFile("msg.txt")

	if err != nil {
		log.Println("Error in Read file")
		return
	}

	fmt.Println(string(content))

	/*
		Read & Write file CSV
		=====================
	*/

}
