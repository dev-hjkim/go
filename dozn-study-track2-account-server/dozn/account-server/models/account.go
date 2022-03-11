package models

import (
	"github.com/google/uuid"
    "gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      	string    	`gorm:"primaryKey"`
	AcctNum 	string  	`json:"accountNumber"`
	AcctTp    	string 		`json:"accountType"`
	AcctNick	string		`json:"accountNickname"`
	Balance		uint		`json:"balance"`
}

func (acct *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	acct.ID = uuid.NewString()
	return
  }
