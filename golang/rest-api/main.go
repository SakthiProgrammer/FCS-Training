package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Restaurant struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   uint   `json:"rating"`
	Type     string `json:"type"`
	Foods    []Food `json:"foods"`
}

type Food struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price uint   `json:"price"`
}

var restaurants []Restaurant

var foods []Food

/*
path := r.URL.Path

	// Split the URL path by "/"localhst:8080/13223
	parts := strings.Split(path, "/")

	// Assuming the structure is "/restaurants/{id}", the ID will be in parts[2]
	if len(parts) >= 3 {
		restaurantID := parts[2] // This is where the path variable is located
		// Process the restaurantID as needed
		fmt.Fprintf(w, "Restaurant ID: %s\n", restaurantID)
	} else {
		http.Error(w, "Invalid path", http.StatusBadRequest)
	}

*/
func main() {
	log.Println("Server Started")
	f, err := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	foods = append(foods, Food{Id: 1, Name: "Biriyani", Type: "Non-veg", Price: 120}, Food{1, "Chicken Fried Rice", "Non-veg", 100})
	log.SetOutput(f)
	restaurants = append(restaurants, Restaurant{Id: 1, Name: "Sakthi Hotal", Location: "Chennai KK nagar", Rating: 4, Type: "Non-veg Only"})
	restaurants = append(restaurants, Restaurant{2, "Ram Restro.", "Tiruvannamalai KK nagar", 2, "Veg-Only", foods})

	http.HandleFunc("/restaurants", GetRestaurant)
	http.HandleFunc("/random", randomImageGenerator)

	// http.Handler("/restaurants/:id", getRestaurantById)
	// http.HandleFunc("/restaurants/add", addRestaurant)
	http.ListenAndServe(":3000", nil)

}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// r.Header = http.Header{
	// 	"Host":          {"www.host.com"},
	// 	"Content-Type":  {"application/json"},
	// 	"Authorization": {"Bearer Token"},
	// }
	log.Println("GR001-(+)")

	fmt.Println(restaurants)

	q := r.URL.RawQuery

	// q1 := r.URL.RawPath

	// fmt.Println(q1)

	fmt.Println(q)

	if r.Method == "GET" {

		data, err := json.Marshal(restaurants)

		if err != nil {
			fmt.Fprintf(w, " Error taking Data "+err.Error())
			return
		}
		(w).Header().Set("id", strconv.Itoa(restaurants[0].Id))
		fmt.Fprintf(w, string(data))

	} else if r.Method == "POST" {
		var restaurant Restaurant

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println("Error occur in ReadAll", err)
			return
		} else {
			err = json.Unmarshal(body, &restaurant)
			if err != nil {
				fmt.Fprintf(w, "Error occur in Unmarshaling"+err.Error())
				return
			}
			data, err := json.Marshal(restaurant)
			if err != nil {
				log.Println("Error occur in Marshal", err)
				return
			}
			restaurants = append(restaurants, restaurant)
			fmt.Fprintf(w, string(data))
		}
	} else if r.Method == "PUT" {
		id := r.URL.Query().Get("id")
		fmt.Println(id)

		intId, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error Occur during converting String to Integer")
		}

		var restaurant Restaurant

		for i := 0; i < len(restaurants); i++ {
			if restaurants[i].Id == intId {
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Error Occur in Read Request ", err.Error())
				} else {
					err = json.Unmarshal(data, &restaurant)
					if err != nil {
						log.Println("Error Occur in Unmarshaling")
					} else {
						restaurants[i].Name = restaurant.Name
						restaurants[i].Location = restaurant.Location
						restaurants[i].Type = restaurant.Type
						fmt.Println(restaurant)
					}
				}
				return
			}
		}

	}
	log.Println("GR001-(-)")
}

func randomImageGenerator(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")
	url := "https://dog.ceo/api/breeds/image/random"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	lBody, lErr := io.ReadAll(res.Body)

	if lErr != nil {
		fmt.Println("Error to Convert to Json to Byte")
		return
	}

	var data map[string]string

	err = json.Unmarshal(lBody, &data)

	if err != nil {
		fmt.Println("Error to Convert Unmarshal")
		return
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error to Convert Marshal")
		return
	}

	fmt.Println(data)
	fmt.Fprintf(w, string(jsonData))
}

// func getRestaurantById(w http.ResponseWriter, r *http.Request){

// }

// func findRestaurantById(id int) int {
// 	for i := 0; i < len(restaurants); i++{
// 		if restaurants[i].Id == id{
// 			return i
// 		}
// 	}
// 	return -1
// }

// func addRestaurant(w http.ResponseWriter, r *http.Request) {
// 	(w).Header().Set("Access-Control-Allow-Methods", "POST")

// 	log.Println("M002-(+)")

// 	if r.Method == http.MethodPost {
// 		var restaurant Restaurant

// 		body, err := ioutil.ReadAll(r.Body)

// 		if err != nil {
// 			log.Println("Error occur in ReadAll", err)
// 			return
// 		} else {
// 			err = json.Unmarshal(body, &restaurant)
// 			if err != nil {
// 				fmt.Fprintf(w, "Error occur in Unmarshaling"+err.Error())
// 				return
// 			}
// 			data, err := json.Marshal(restaurant)
// 			if err != nil {
// 				log.Println("Error occur in Marshal", err)
// 				return
// 			}
// 			restaurants = append(restaurants, restaurant)
// 			fmt.Fprintf(w, string(data))

// 		}
// 		log.Println("M002-(-)")

// 	}
// }
