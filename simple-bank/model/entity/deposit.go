package entity

import "time"

type Deposit struct {
	Deposit_Code   string    `json:"deposit_code"`
	User_Id        int       `json:"user_id"`
	Deposit_Amount float32   `json:"deposit_amount"`
	Description    string    `json:"description"`
	Date_Time      time.Time `json:"date_time"`
}

type DepositRequest struct {
	Deposit_Amount float32 `json:"deposit_amount"`
	Description    string  `json:"description"`
}
