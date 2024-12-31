package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	res "user/apps"
	e "user/common"
)

var UserData []res.User

/*=======================================================Get All User=======================================================*/

func GetUsers(w http.ResponseWriter, r *http.Request) {

	log.Println("GetUsers-(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var lUserRes res.UserRes

	lUser := res.User{1, "sakthi", "sakthiveld@gmail.com", 9047684797}

	UserData = append(UserData, lUser, res.User{2, "ram", "ram@gmail.com", 1231312313})
	lUserRes.Status = e.SUCCESS_CODE

	lIdStr := r.Header.Get("ID")

	if r.Method != http.MethodGet {
		lUserRes.Status = e.ERROR_CODE
		lUserRes.ErrMsg = "Method Not Allowed,GET Method Only  Allowed"

	} else {
		// lData, lErr := json.Marshal(lUserRes)
		// if lErr != nil {

		// 	log.Println("MGU-001: ", lErr.Error())
		// 	log.Println(lErr)
		// 	lUserRes.ErrMsg = lErr.Error() + "MGU-001"
		// 	lUserRes.Status = e.ERROR_CODE
		// } else {
		//
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// fmt.Fprintf(w, string(lData))

		if lIdStr == "" {

		} else {

		}

		lUserRes.UserArr = UserData
		lUserRes.ErrMsg = ""
	}

	// fmt.Fprintf(w, "Works Correct")

	lData, lErr := json.Marshal(lUserRes)

	if lErr != nil {

		log.Println("MGU-004: ", lErr.Error())
		lUserRes.ErrMsg = lErr.Error()
		lUserRes.Status = e.ERROR_CODE
		fmt.Fprintf(w, string(lData))
	} else {

		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetUsers-(-)")
}

/*=======================================================Get User By ID=======================================================*/

func GetUser(w http.ResponseWriter, r *http.Request) {

	log.Println("GetUser-(+)")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	lIdStr := r.Header.Get("ID")

	var lUserRes res.UserRes
	lIntId, lerr := strconv.Atoi(lIdStr)

	if r.Method != http.MethodGet {

		lUserRes.Status = e.ERROR_CODE
		lUserRes.ErrMsg = "Method Not Allowed,GET Method Only  Allowed"
		// lData, lErr := json.Marshal(lUserRes)

		// if lErr != nil {
		// 	log.Println("Error-02: ", lErr.Error())
		// 	log.Println(lErr)
		// 	lUserRes.ErrMsg = lErr.Error()
		// 	lUserRes.Status = e.ERROR_CODE
		// } else {

		// 	w.WriteHeader(http.StatusMethodNotAllowed)
		// 	fmt.Fprintf(w, string(lData))
		// }
		// return
	} else if lIdStr == "" {

		lUserRes.ErrMsg = "Provide a value for id"
		lUserRes.Status = e.ERROR_CODE
		// w.WriteHeader(http.StatusBadRequest)

		// lData, lErr := json.Marshal(lUserRes)

		// if lErr != nil {
		// 	log.Println("Error-01", lErr.Error())
		// } else {
		// 	fmt.Fprintf(w, string(lData))
		// }
	} else if lerr != nil {

		lUserRes.ErrMsg = "Provide a Value in as a Number"
		lUserRes.Status = e.ERROR_CODE
		// w.WriteHeader(http.StatusBadRequest)

		// lData, lErr := json.Marshal(lUserRes)

		// if lErr != nil {
		// 	log.Println("Error-02", lErr.Error())
		// } else {
		// 	fmt.Fprintf(w, string(lData))
		// }
	} else if !(lIntId >= 1 && lIntId <= 100) {

		lUserRes.ErrMsg = "Provide a Between 1 to 100"
		lUserRes.Status = e.ERROR_CODE
		w.WriteHeader(http.StatusBadRequest)

		// lData, lErr := json.Marshal(lUserRes)

		// if lErr != nil {
		// 	log.Println("Error-02", lErr.Error())
		// } else {
		// 	fmt.Fprintf(w, string(lData))
		// }
		// return
	} else {
		fmt.Println(UserData)

		var lUser res.User

		for i := 0; i < len(UserData); i++ {
			if UserData[i].ID == lIntId {
				lUser = UserData[i]
				lUserRes.Status = e.SUCCESS_CODE
				lUserRes.ErrMsg = "Success"
				lUserRes.UserArr = append(lUserRes.UserArr, lUser)
				break
				// lData, lErr := json.Marshal(lUserRes)
				// if lErr != nil {
				// 	log.Println("Error-03", lErr.Error())
				// 	w.WriteHeader(http.StatusBadRequest)
				// } else {
				// 	fmt.Fprintf(w, string(lData))
				// }
			}
		}

		if lUser == (res.User{}) {
			lUserRes.ErrMsg = "User with Id: " + lIdStr + " Not Found"
			lUserRes.Status = e.ERROR_CODE
		}
	}

	lData, lErr := json.Marshal(lUserRes)

	if lErr != nil {
		log.Println("Error-04", lErr.Error())
	} else {
		fmt.Fprintf(w, string(lData))
	}
	// w.WriteHeader(http.StatusNotFound)

	log.Println("GetUser-(-)")
}

/*=======================================================Insert User=======================================================*/

func AddUser(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, PUT, OPTIONS")

	// lMethod := r.Header.Get("TYPE")

	log.Println("M003-(+)")

	var lUserRes res.UserRes

	if r.Method != http.MethodPost {
		lUserRes.Status = e.ERROR_CODE
		lUserRes.ErrMsg = "Method Not Allowed,GET Method Only  Allowed"

	} else {

		var lUser res.User

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("Error-01: ", lErr.Error())
			log.Println(lErr)
			lUserRes.ErrMsg = lErr.Error()
			lUserRes.Status = e.ERROR_CODE

		}

		lErr = json.Unmarshal(lData, &lUser)

		if lErr != nil {

			log.Println("MG: ", lErr.Error())
			log.Println(lErr)
			lUserRes.ErrMsg = lErr.Error()
			lUserRes.Status = e.ERROR_CODE
		} else {
			UserData = append(UserData, lUser)
			lUserRes.UserArr = UserData
			lUserRes.Status = e.SUCCESS_CODE
			lUserRes.ErrMsg = ""
		}

		// idEmpty

		log.Println("M003-(-)")

	}
}

// func isEmpty(val string) bool{

// 	if val == ""{
// 		return false
// 	}

// }
