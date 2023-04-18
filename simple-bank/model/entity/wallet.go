package entity

type Wallet struct {
	Wallet_Id int     `json:"wallet_id"`
	User_Id   int     `json:"user_id"`
	Amount    float32 `json:"amount"`
}
