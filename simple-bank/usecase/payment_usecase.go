package usecase

import (
	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/repository"
)

type PaymentUsecase interface {
	TransferBalance(newPayment *entity.PaymentInfo, senderId int) (entity.Payment, error)
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}

func (u *paymentUsecase) TransferBalance(newPayment *entity.PaymentInfo, senderId int) (entity.Payment, error) {
	return u.paymentRepo.CreatePayment(newPayment, senderId)
}
