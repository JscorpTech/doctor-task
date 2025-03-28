package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorID uint      `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	Time     time.Time `gorm:"unique;not null"`
	Status   string    `gorm:"default:'scheduled'"` // "scheduled", "cancelled"
	User     User
	Doctor   Doctor
}

func (d *Appointment) TableName() string {
	return "appointment"
}
