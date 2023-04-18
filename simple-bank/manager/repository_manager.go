package manager

import (
	"github.com/rizkyfazri23/mnc/repository"
	"gorm.io/gorm"
)

type RepoManager interface {
	UserRepo() repository.UserRepo
	PaymentRepo() repository.PaymentRepository
	DepositRepo() repository.DepositRepository
	HistoryRepo() repository.HistoryRepository
}

type repositoryManager struct {
	db *gorm.DB
}

func (r *repositoryManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepository(r.db)
}

func (r *repositoryManager) PaymentRepo() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.db)
}

func (r *repositoryManager) DepositRepo() repository.DepositRepository {
	return repository.NewDepositRepository(r.db)
}

func (r *repositoryManager) HistoryRepo() repository.HistoryRepository {
	return repository.NewHistoryRepository(r.db)
}

func NewRepoManager(db *gorm.DB) RepoManager {
	return &repositoryManager{
		db: db,
	}
}
