package entity

type User struct {
	User_Id  int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Amount   int    `json:"amount"`
}
