package repository

import (
	"github.com/JscorpTech/jst-go/models"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{
		DB: db,
	}
}

func (a *AppointmentRepository) List() ([]*models.Appointment, error) {
	var doctors []*models.Appointment
	if err := a.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
