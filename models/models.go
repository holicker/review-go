package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Review struct{
	gorm.Model
	Id int `json:"id"`
	VendorId int `json:"vendorId"`
	Writer string `json:"writer"`
	Content string `json:"content"`
	RegisterdDate time.Time `json:"registerdDate"`
}

func (Review) TableNale() string {
	return "review"
}