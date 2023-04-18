package entity

type Member struct {
	Member_Id      int     `json:"member_id"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Email_Address  string  `json:"email_address"`
	Contact_Number string  `json:"contact_number"`
	Wallet_Amount  float64 `json:"wallet_amount"`
	Status         bool    `json:"status"`
}

type MemberLogin struct {
	Member_Id int    `json:"member_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
