package entity

import "time"

type Token struct {
	Token     string    `json:"token"`
	Expire_At time.Time `json:"expire_at"`
	Status    bool      `json:"status"`
	User_Id   int       `json:"user_id"`
}
