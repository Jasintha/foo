package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

func (Address) TableName() string {
	return "address"
}
func (m *Address) PreloadAddress(db *gorm.DB) *gorm.DB {
	return db
}
