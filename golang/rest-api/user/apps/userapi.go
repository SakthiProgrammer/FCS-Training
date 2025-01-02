package apps

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phonenumber"`
}

type UserRes struct {
	UserArr []User `json:"usersArr"`
	ErrMsg  string `json:"message"`
	Status  string `json:"status"`
}
