package main

import (
	"log"
	"net/http"
	"os"
	"time"
	m "user/methods"
)

func main() {
	log.Println("Server Started")
	f, err := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/users", m.GetUsers)
	// http.HandleFunc("/user", m.GetUser)
	http.HandleFunc("/adduser", m.AddUser)
	http.HandleFunc("/updateuser", m.UpdateUser)
	http.HandleFunc("/tapi", m.GetGendreize)
	http.ListenAndServe(":3000", nil)
}
