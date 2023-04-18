package repository

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/rizkyfazri23/mnc/model/entity"
)

type PaymentRepository interface {
	CreatePayment(newPayment *entity.PaymentInfo, senderId int) (entity.Payment, error)
	CheckToken(userId int) error
}

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func (r *PaymentRepositoryImpl) CheckToken(userId int) error {

	var t entity.Token
	err := r.db.Where("user_id = ? AND status = true", userId).First(&t).Error
	if err != nil {
		return errors.New("Session Expired")
	}
	return nil
}

func (r *PaymentRepositoryImpl) CreatePayment(newPayment *entity.PaymentInfo, senderId int) (entity.Payment, error) {
	err := r.CheckToken(senderId)

	if err != nil {
		return entity.Payment{}, err
	}

	var receiptId int
	if err := r.db.Model(&entity.User{}).Where("username = ?", newPayment.ReceiptUsername).Select("user_id").Take(&receiptId).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("Get ReceiptId")
	}

	if senderId == receiptId {
		err := errors.New("can't transfer to yourself")
		log.Println("Sender and recipient usernames are identical")
		return entity.Payment{}, err
	} else {
		log.Println("Diff username")
	}

	var senderBalance float32
	if err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", senderId).Select("amount").Take(&senderBalance).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("Get SenderBalance")
	}

	if senderBalance < newPayment.Payment_Amount {
		err := errors.New("you don't have that much money")
		log.Println("Insufficient funds")
		return entity.Payment{}, err
	} else {
		log.Println("Sufficient funds")
	}

	var payment entity.Payment
	if err := r.db.Raw("INSERT INTO payments (sender_id, receipt_id, payment_amount, description, date_time) "+
		"VALUES (?, ?, ?, ?, ?) RETURNING payment_id, payment_code, sender_id, receipt_id, payment_amount, description, date_time",
		senderId, receiptId, newPayment.Payment_Amount, newPayment.Description, time.Now()).Scan(&payment).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("Payment Created")
	}

	if err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", senderId).Update("amount", gorm.Expr("amount - ?", newPayment.Payment_Amount)).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("funds transferred")
	}

	if err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", receiptId).Update("amount", gorm.Expr("amount + ?", newPayment.Payment_Amount)).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("funds received")
	}

	transactionLog := &entity.TransactionLog{
		UserID:          payment.Sender_Id,
		TransactionCode: payment.Payment_Code,
		TransactionType: "Payment",
		Amount:          payment.Payment_Amount,
	}

	if err := r.db.Model(&entity.TransactionLog{}).Omit("id").Create(transactionLog).Error; err != nil {
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("Transaction log created")
	}

	return payment, nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}
