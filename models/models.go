package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	VendorId      int       `gorm:"column:vendor_id" json:"vendorId"`
	Writer        string    `gorm:"column:writer" json:"writer"`
	Title         string    `gorm:"column:title" json:"title"`
	Content       string    `gorm:"column:content" json:"content"`
	RegisterdDate time.Time `gorm:"column:registerd_date" json:"registerdDate"`
}

func (Review) TableNale() string {
	return "review"
}
