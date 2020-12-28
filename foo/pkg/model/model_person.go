package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	Model
	Name      string  `json:"name"`
	Address   Address `json:"address"`
	AddressID int64
}

func (Person) TableName() string {
	return "person"
}
func (m *Person) PreloadPerson(db *gorm.DB) *gorm.DB {
	return db.Preload("Address")
}
