package repository

import (
	"time"

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

func (a *AppointmentRepository) List() ([]models.Appointment, error) {
	var doctors []models.Appointment
	if err := a.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func (a *AppointmentRepository) FindByDoctor(id uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	if err := a.DB.Where("doctor_id = ?", id).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (a *AppointmentRepository) DoctorAppointmentBetween(doctor uint, start, end time.Time) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := a.DB.Where("time >= ? AND  time <= ?", start, end).Where("doctor_id = ?", doctor).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (a *AppointmentRepository) Create(user_id, doctor uint, time time.Time, status string) (*models.Appointment, error) {
	user := &models.Appointment{
		UserID:   user_id,
		DoctorID: doctor,
		Time:     time,
		Status:   status,
	}
	if err := a.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
