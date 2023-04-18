package usecase

import (
	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/repository"
)

type DepositUsecase interface {
	Add(newDeposit *entity.DepositRequest, user_id int) (entity.Deposit, error)
}

type depositUsecase struct {
	depositRepo repository.DepositRepository
}

func NewDepositUsecase(depositRepo repository.DepositRepository) DepositUsecase {
	return &depositUsecase{
		depositRepo: depositRepo,
	}
}

func (u *depositUsecase) Add(newDeposit *entity.DepositRequest, user_id int) (entity.Deposit, error) {
	return u.depositRepo.MakeDeposit(newDeposit, user_id)
}
