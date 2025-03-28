package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Phone     string    `gorm:"unique"`
	Specialty string    `gorm:"not null"`
	WorkStart time.Time `gorm:"not null;type:time"`
	WorkEnd   time.Time `gorm:"not null;type:time"`
}

func (d *Doctor) TableName() string {
	return "doctor"
}
