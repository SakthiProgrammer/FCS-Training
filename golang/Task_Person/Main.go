package main

import (
	"fmt"
)

//Person type
type Person struct {
	Id      int
	Name    string
	Handsup bool
	Gender  string
	Pair
}

//to Store List of Person
var persons []Person

type Pair struct {
	Id int
}

// person id : pair person id
var mapPair = make(map[int]int)

func main() {

	again := true
	for again {
		var name, gender string
		var pairId int
		fmt.Printf("Enter Name: ")
		fmt.Scan(&name)
		fmt.Printf("Enter Gender(m/f): ")
		fmt.Scan(&gender)

		id := len(persons) + 1

		fmt.Printf("All the Person : \n    %+v   \n", persons)
		fmt.Printf("Enter the Id of the Person do you want to Pair (0 - for no pair): ")
		fmt.Scan(&pairId)

		//person id : pair person id
		var pair Pair
		pair.Id = pairId
		if addNewPair(id, pairId) {
			fmt.Println("Pair Successfully")
		} else {
			if len(mapPair) == 0 {
				mapPair[id] = pairId
			} else {
				mapPair[id] = 0
				pair.Id = 0
				fmt.Println("unable to Pair He is Already Pair with Other Person")
			}
		}

		fmt.Printf("Enter the Person Id for Hands Up: ")
		var idForHandsUp int
		fmt.Scan(&idForHandsUp)

		toHandsUp(idForHandsUp)

		newPerson := Person{Id: id, Name: name, Gender: gender, Pair: pair}
		fmt.Println(newPerson)

		persons = append(persons, newPerson)
		fmt.Printf("Updated Person: \n    %v   \n", persons)

		// ================================================= Questions ===================================================

		count := 0
		numberOfPersonsHandsUp(&count)
		fmt.Println("Q1: Number of Person Hands are Up: ", count)
		fmt.Println("Q2: Number of Pair Person Hands are Up: ", NumberOfPairsAreHandsUp())
		maleHandsUpCount, femaleHandsUpCount := NumberofNonPairPersonhandsUp()
		fmt.Println("Q3: Number of Non Pair Person hands are Up: ", len(persons)-NumberOfPairsAreHandsUp())
		fmt.Printf("Q4: Total Hands Up person Gender Male: %v | Female: %v", maleHandsUpCount, femaleHandsUpCount)
		fmt.Println("\n---------Pairs----------")
		displayListOfPair()

		// =================================================================================================================

		for key, value := range mapPair {
			fmt.Println(key, " : ", value)
		}
	}
}

func toHandsUp(id int) {
	for i := 0; i < len(persons); i++ {
		if persons[i].Id == id {
			persons[i].Handsup = true
			fmt.Printf("Successfully Handsup")
			break
		}
	}
}

func numberOfPersonsHandsUp(count *int) {
	for i := 0; i < len(persons); i++ {
		if persons[i].Handsup {
			*count++
		}
	}
}

func displayListOfPair() {
	for key, value := range mapPair {
		if persons[key-1].Id == key && value != 0 {
			fmt.Println(persons[key-1].Name, " : ", persons[value-1].Name)
		}
	}
}

func addNewPair(id, pairId int) bool {
	for key, value := range mapPair {
		if key == pairId && value == 0 {
			delete(mapPair, pairId)
			mapPair[id] = pairId
			for j := 0; j < len(persons); j++ {
				if persons[j].Id == pairId {
					persons[j].Pair.Id = id
					break
				}
			}
			return true
		}
	}
	return false
}

// func isAlreadyPair(pairId int) bool{

// 	for
// }
// func totalHandsUpGender() int {

// }

func NumberOfPairsAreHandsUp() int {
	pairCount := 0
	for key, value := range mapPair {
		if persons[key-1].Pair.Id == value && persons[value].Id == key && persons[key-1].Handsup && persons[value].Handsup {
			pairCount++
		}
	}
	return pairCount
}

func NumberofNonPairPersonhandsUp() (int, int) {
	maleHandsUpCount := 0
	femaleHandsUpCount := 0
	for i := 0; i < len(persons); i++ {
		if persons[i].Handsup {
			if persons[i].Gender == "m" {
				maleHandsUpCount++
			} else if persons[i].Gender == "f" {
				femaleHandsUpCount++
			}
		}
	}
	return maleHandsUpCount, femaleHandsUpCount
}
