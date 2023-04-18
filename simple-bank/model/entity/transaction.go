package entity

type TransactionLog struct {
	ID              int     `gorm:"column:id;primaryKey"`
	UserID          int     `gorm:"column:user_id"`
	TransactionCode string  `gorm:"column:transaction_code"`
	TransactionType string  `gorm:"column:transaction_type"`
	Amount          float32 `gorm:"column:amount"`
}
