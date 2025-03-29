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

func (d *DoctorUsecase) List(search string) []*domain.DoctorResponse {
	var doctorsDTO []*domain.DoctorResponse
	doctors, _ := d.DoctorRepository.List(search)
	for _, e := range doctors {
		doctorsDTO = append(doctorsDTO, &domain.DoctorResponse{
			ID:        e.ID,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			Phone:     e.Phone,
			Specialty: e.Specialty,
			WorkStart: e.WorkStart.Format("15:04"),
			WorkEnd:   e.WorkEnd.Format("15:04"),
		})
	}
	return doctorsDTO
}

func (d *DoctorUsecase) FindByID(id uint) (*domain.DoctorDetailResponse, error) {
	doctor, err := d.DoctorRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	var appointmentsDTO []*domain.AppointmentResponse
	for _, appointment := range doctor.Appointments {
		appointmentsDTO = append(appointmentsDTO, &domain.AppointmentResponse{
			ID:     appointment.ID,
			Time:   appointment.Time,
			Status: appointment.Status,
		})
	}

	return &domain.DoctorDetailResponse{
		DoctorResponse: domain.DoctorResponse{
			ID:        doctor.ID,
			FirstName: doctor.FirstName,
			LastName:  doctor.LastName,
			Phone:     doctor.Phone,
			Specialty: doctor.Specialty,
			WorkStart: doctor.WorkStart.Format("15:04"),
			WorkEnd:   doctor.WorkEnd.Format("15:04"),
		},
		Appointments: appointmentsDTO,
	}, nil
}
