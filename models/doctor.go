package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Specialty string `gorm:"not null"`
	WorkHours string `gorm:"not null"`
}

func (d *Doctor) TableName() string {
	return "doctor"
}
