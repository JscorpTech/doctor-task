package repository

import (
	"github.com/JscorpTech/jst-go/models"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	DB *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{
		DB: db,
	}
}

func (d *DoctorRepository) List() ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	if err := d.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
