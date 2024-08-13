package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

type Task struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"default:'Todo'"`
	Priority    string    `gorm:"default:'Medium'"`
	DueDate     time.Time
	UserID      uint      `gorm:"not null"`
}
