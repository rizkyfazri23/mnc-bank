package usecase

import (
	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/repository"
)

type HistoryUsecase interface {
	GetAllTransaction(userID int) ([]entity.TransactionLog, error)
	GetAllAuth(userID int) ([]entity.AuthenticationLog, error)
}

type historyUsecase struct {
	historyRepo repository.HistoryRepository
}

func NewHistoryUsecase(historyRepo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{
		historyRepo: historyRepo,
	}
}

func (u *historyUsecase) GetAllTransaction(userID int) ([]entity.TransactionLog, error) {
	return u.historyRepo.TransactionHistory(userID)
}

func (u *historyUsecase) GetAllAuth(userID int) ([]entity.AuthenticationLog, error) {
	return u.historyRepo.AuthHistory(userID)
}
