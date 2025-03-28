package usecase

import (
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/repository"
)

type AppointmentUsecase struct {
	AppointmentRepository repository.AppointmentRepository
}

func NewAppointmentUsecase(appointmentRepository repository.AppointmentRepository) domain.AppointmentUsecase {
	return &AppointmentUsecase{
		AppointmentRepository: appointmentRepository,
	}
}

func (d *AppointmentUsecase) List() ([]*models.Appointment, error) {
	appointments, err := d.AppointmentRepository.List()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
