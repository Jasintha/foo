package model

import "time"

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Tabler interface {
	TableName() string
}

func InitModels(db *gorm.DB) {

	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Person{})
}
