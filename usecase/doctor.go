package usecase

import (
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/repository"
)

type DoctorUsecase struct {
	DoctorRepository repository.DoctorRepository
}

func NewDoctorUsecase(doctorRepository repository.DoctorRepository) domain.DoctorUsecase {
	return &DoctorUsecase{
		DoctorRepository: doctorRepository,
	}
}

func (d *DoctorUsecase) List() ([]*models.Doctor, error) {
	doctors, err := d.DoctorRepository.List()
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
