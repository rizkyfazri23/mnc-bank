package repository

import (
	"github.com/rizkyfazri23/mnc/model/entity"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	TransactionHistory(userID int) ([]entity.TransactionLog, error)
	AuthHistory(userID int) ([]entity.AuthenticationLog, error)
}

type historyRepository struct {
	db *gorm.DB
}

func (r *historyRepository) TransactionHistory(userID int) ([]entity.TransactionLog, error) {
	var logs []entity.TransactionLog
	err := r.db.Where("user_id = ?", userID).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *historyRepository) AuthHistory(userID int) ([]entity.AuthenticationLog, error) {
	var logs []entity.AuthenticationLog
	err := r.db.Where("user_id = ?", userID).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	repo := new(historyRepository)
	repo.db = db
	return repo
}
