package manager

import "github.com/rizkyfazri23/mnc/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	PaymentUsecase() usecase.PaymentUsecase
	DepositUsecase() usecase.DepositUsecase
	HistoryUsecase() usecase.HistoryUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func (u *usecaseManager) PaymentUsecase() usecase.PaymentUsecase {
	return usecase.NewPaymentUsecase(u.repoManager.PaymentRepo())
}

func (u *usecaseManager) DepositUsecase() usecase.DepositUsecase {
	return usecase.NewDepositUsecase(u.repoManager.DepositRepo())
}

func (u *usecaseManager) HistoryUsecase() usecase.HistoryUsecase {
	return usecase.NewHistoryUsecase(u.repoManager.HistoryRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
