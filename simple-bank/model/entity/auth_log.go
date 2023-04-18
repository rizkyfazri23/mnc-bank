package entity

type AuthenticationLog struct {
	ID     int    `gorm:"primaryKey"`
	UserID int    `gorm:"not null"`
	Status string `gorm:"not null"`
}
