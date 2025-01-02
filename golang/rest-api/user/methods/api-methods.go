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

	// lUser := res.User{1, "sakthi", "sakthiveld@gmail.com", 9047684797}

	// UserData = append(UserData, lUser, res.User{2, "ram", "ram@gmail.com", 1231312313})
	lUserRes.Status = e.SUCCESS_CODE

	lIdStr := r.Header.Get("ID")

	if r.Method != http.MethodGet {
		lUserRes.Status = e.ERROR_CODE
		lUserRes.ErrMsg = "Method Not Allowed,GET Method Only  Allowed hello"

	} else {

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

	} else if lIdStr == "" {

		lUserRes.ErrMsg = "Provide a value for id"
		lUserRes.Status = e.ERROR_CODE

	} else if lerr != nil {

		lUserRes.ErrMsg = "Provide a Value in as a Number"
		lUserRes.Status = e.ERROR_CODE

	} else if !(lIntId >= 1 && lIntId <= 100) {

		lUserRes.ErrMsg = "Provide a Between 1 to 100"
		lUserRes.Status = e.ERROR_CODE
		w.WriteHeader(http.StatusBadRequest)

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
			}
		}

		if lUser == (res.User{}) {
			lUserRes.ErrMsg = "User with Id: " + lIdStr + " Not Found"
			lUserRes.Status = e.ERROR_CODE
		}
	}

	lData, lErr := json.Marshal(lUserRes)

	if lErr != nil {
		log.Println("Error-004", lErr.Error())
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetUser-(-)")
}

/*=======================================================Insert User=======================================================*/

func AddUser(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, PUT, OPTIONS")

	// lMethod := r.Header.Get("TYPE")

	log.Println("AddUser-(+)")

	var lUserRes res.UserRes
	lUserRes.Status = e.SUCCESS_CODE

	if r.Method != http.MethodPost {
		lUserRes.Status = e.ERROR_CODE
		lUserRes.ErrMsg = "Method Not Allowed,GET Method Only  Allowed"

	} else {

		var lUser res.User
		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("MGU-005: ", lErr.Error())
			log.Println(lErr)
			lUserRes.ErrMsg = lErr.Error()
			lUserRes.Status = e.ERROR_CODE
		}

		lErr = json.Unmarshal(lData, &lUser)

		if lErr != nil {
			log.Println("MGU-006: ", lErr.Error())
			log.Println(lErr)
			lUserRes.ErrMsg = lErr.Error()
			lUserRes.Status = e.ERROR_CODE
		} else {

			if lUser.ID == 0 {
				lUserRes.Status = e.ERROR_CODE
				lUserRes.ErrMsg = "ERR: id field is empty"
			} else if lUser.Name == "" {
				lUserRes.Status = e.ERROR_CODE
				lUserRes.ErrMsg = "ERR: name field is empty"
			} else if lUser.Email == "" {
				lUserRes.Status = e.ERROR_CODE
				lUserRes.ErrMsg = "ERR: email field is empty"
			} else if strconv.Itoa(lUser.PhoneNumber) == "0" {
				lUserRes.Status = e.ERROR_CODE
				lUserRes.ErrMsg = "ERR: phonenumber field is empty"
			} else if 10 != len(strconv.Itoa(lUser.PhoneNumber)) {
				lUserRes.Status = e.ERROR_CODE
				lUserRes.ErrMsg = "ERR: Invalid phonenumber, must be 10 digit field is empty"

			} else {
				UserData = append(UserData, lUser)
				lUserRes.UserArr = []res.User{UserData[len(UserData)-1]}
				lUserRes.Status = e.SUCCESS_CODE
				lUserRes.ErrMsg = ""
			}

		}

		//marshal
		lData, lErr = json.Marshal(lUserRes)

		if lErr != nil {
			log.Println("Error-04", lErr.Error())
		} else {
			fmt.Fprintf(w, string(lData))
		}

		log.Println("AddUser-(-)")
	}
}

func Genderize(w http.ResponseWriter, r *http.Request) {

}

//https://api.genderize.io?name=sakthivel

/*{
    "count": 1949,
    "name": "sakthivel",
    "gender": "male",
    "probability": 0.96
}
*/
