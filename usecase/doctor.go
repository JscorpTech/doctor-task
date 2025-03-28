package usecase

import (
	"github.com/JscorpTech/jst-go/domain"
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

func (d *DoctorUsecase) List() ([]*domain.DoctorResponse, error) {
	var doctorsDTO []*domain.DoctorResponse
	doctors, err := d.DoctorRepository.List()
	for _, e := range doctors {
		doctorsDTO = append(doctorsDTO, &domain.DoctorResponse{
			FirstName: e.FirstName,
			LastName:  e.LastName,
			Phone:     e.Phone,
			Specialty: e.Specialty,
			WorkStart: e.WorkStart.Format("15:04"),
			WorkEnd:   e.WorkEnd.Format("15:04"),
		})
	}
	if err != nil {
		return nil, err
	}
	return doctorsDTO, nil
}
