package repository

import (
	"fmt"

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

func (d *DoctorRepository) List(search string) ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	search = fmt.Sprintf("%%%s%%", search)

	if err := d.DB.Where("first_name LIKE ? OR last_name LIKE ? OR phone LIKE ?",
		search, search, search).
		Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func (d *DoctorRepository) FindByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := d.DB.Preload("Appointments").First(&doctor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}
