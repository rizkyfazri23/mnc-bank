package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindOne(id int) (entity.User, error)
	Create(newUser *entity.User) (entity.User, error)
	Update(user *entity.User, userId int) (entity.User, error)
	LoginCheck(username string, password string) (string, error)
	Logout(userId int) (string, error)
	CheckToken(userId int) error
}

type userRepo struct {
	db *gorm.DB
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *userRepo) CheckToken(userId int) error {
	var t entity.Token
	err := r.db.Where("user_id = ? AND status = true", userId).First(&t).Error
	if err != nil {
		return errors.New("Session Expired")
	}
	return nil
}

func (r *userRepo) LoginCheck(username string, password string) (string, error) {
	var u entity.User
	err := r.db.Where("username = ?", username).First(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.User_Id)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	tokenData := entity.Token{
		Token:     token,
		User_Id:   u.User_Id,
		Status:    true,
		Expire_At: expirationTime,
	}
	if err := r.db.Create(&tokenData).Error; err != nil {
		return "", err
	}

	authLog := entity.AuthenticationLog{
		UserID: u.User_Id,
		Status: "Login",
	}
	if err := r.db.Create(&authLog).Error; err != nil {
		return "", err
	}

	return token, nil
}

func (r *userRepo) FindOne(id int) (entity.User, error) {
	err := r.CheckToken(id)
	if err != nil {
		return entity.User{}, err
	}

	var userInDb entity.User
	err = r.db.Joins("JOIN wallets ON users.user_id = wallets.user_id").
		Select("users.*, wallets.amount").
		Where("users.user_id = ?", id).
		First(&userInDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return entity.User{}, err
	}

	return userInDb, nil
}

func (r *userRepo) Create(newUser *entity.User) (entity.User, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	newUser.Password = string(hashedPassword)

	err = tx.Model(&entity.User{}).Omit("user_id", "amount").Create(newUser).Error
	if err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	var lastUser entity.User
	err = tx.Last(&lastUser).Error
	if err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	wallet := entity.Wallet{
		User_Id: lastUser.User_Id,
		Amount:  0,
	}
	err = tx.Model(&entity.Wallet{}).Omit("wallet_id").Create(&wallet).Error
	if err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	return *newUser, nil
}

func (r *userRepo) Update(user *entity.User, userId int) (entity.User, error) {
	err := r.CheckToken(userId)

	if err != nil {
		return entity.User{}, err
	}

	var userInDb entity.User
	err = r.db.Where("user_id = ?", userId).First(&userInDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, fmt.Errorf("user with id %d not found", userId)
		}
		return entity.User{}, err
	}

	userInDb.Username = user.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	userInDb.Password = string(hashedPassword)

	err = r.db.Where("user_id = ?", userId).Omit("amount").Save(&userInDb).Error
	if err != nil {
		return entity.User{}, err
	}

	return userInDb, nil
}

func (r *userRepo) Logout(userId int) (string, error) {
	var token entity.Token
	err := r.db.Where("user_id = ? AND status = true", userId).First(&token).Error
	if err != nil {
		return "", err
	}

	if err := r.db.Model(&entity.Token{}).Where("token = ?", token.Token).Update("status", false).Error; err != nil {
		return "", err
	}

	log := entity.AuthenticationLog{
		UserID: userId,
		Status: "Logout",
	}
	if err := r.db.Create(&log).Error; err != nil {
		return "", err
	}

	return "Logout success", nil
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{db}
}
