package entity

import "time"

type Payment struct {
	Payment_id     int       `json:"payment_id"`
	Payment_Code   string    `json:"payment_code"`
	Sender_Id      int       `json:"sender_id"`
	Receipt_Id     int       `json:"receipt_id"`
	Payment_Amount float32   `json:"payment_amount"`
	Description    string    `json:"description"`
	Date_time      time.Time `json:"date_time"`
}

type PaymentInfo struct {
	ReceiptUsername string  `json:"receipt_username"`
	Payment_Amount  float32 `json:"payment_amount"`
	Description     string  `json:"description"`
}
