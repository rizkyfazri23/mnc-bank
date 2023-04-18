package repository

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/rizkyfazri23/mnc/model/entity"
)

type DepositRepository interface {
	MakeDeposit(newDeposit *entity.DepositRequest, user_id int) (entity.Deposit, error)
	CheckToken(userId int) error
}

type DepositRepositoryImpl struct {
	db *gorm.DB
}

func (r *DepositRepositoryImpl) CheckToken(userId int) error {

	var t entity.Token
	err := r.db.Where("user_id = ? AND status = true", userId).First(&t).Error
	if err != nil {
		return errors.New("Session Expired")
	}
	return nil
}

func (r *DepositRepositoryImpl) MakeDeposit(newDeposit *entity.DepositRequest, user_id int) (entity.Deposit, error) {
	err := r.CheckToken(user_id)

	if err != nil {
		return entity.Deposit{}, err
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		return entity.Deposit{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Panic: %v\n", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Println("Transaction rolled back")
		} else {
			log.Println("Transaction committed")
			tx.Commit()
		}
	}()

	var deposit entity.Deposit

	if err := tx.Raw("INSERT INTO deposits (user_id, deposit_amount, description, date_time) "+
		"VALUES (?, ?, ?, ?) RETURNING deposit_id, deposit_code, user_id, deposit_amount, description, date_time",
		user_id, newDeposit.Deposit_Amount, newDeposit.Description, time.Now()).Scan(&deposit).Error; err != nil {
		log.Println(err)
		return entity.Deposit{}, err
	} else {
		log.Println("Deposit Created")
		deposit = entity.Deposit{
			User_Id:        user_id,
			Deposit_Amount: newDeposit.Deposit_Amount,
			Description:    newDeposit.Description,
			Date_Time:      time.Now(),
		}
	}

	if err := tx.Model(&entity.Wallet{}).Where("user_id = ?", user_id).Update("amount", gorm.Expr("amount + ?", newDeposit.Deposit_Amount)).Error; err != nil {
		log.Println(err)
		return entity.Deposit{}, err
	} else {
		log.Println("funds transferred")
	}

	// create transaction log
	transactionLog := &entity.TransactionLog{
		UserID:          user_id,
		TransactionCode: deposit.Deposit_Code,
		TransactionType: "Deposit",
		Amount:          deposit.Deposit_Amount,
	}

	if err := tx.Model(&entity.TransactionLog{}).Omit("id").Create(transactionLog).Error; err != nil {
		log.Println(err)
		return entity.Deposit{}, err
	} else {
		log.Println("Transaction log created")
	}

	return deposit, nil
}

func NewDepositRepository(db *gorm.DB) DepositRepository {
	return &DepositRepositoryImpl{db: db}
}
