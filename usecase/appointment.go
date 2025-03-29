package usecase

import (
	"errors"
	"time"

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

func (d *AppointmentUsecase) List() ([]models.Appointment, error) {
	appointments, err := d.AppointmentRepository.List()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (a *AppointmentUsecase) FindByDoctor(id uint) ([]*domain.AppointmentResponse, error) {
	appointments, err := a.AppointmentRepository.FindByDoctor(id)
	var appointmentDTO []*domain.AppointmentResponse
	if err != nil {
		return nil, err
	}
	for _, appointment := range appointments {
		appointmentDTO = append(appointmentDTO, &domain.AppointmentResponse{
			ID:     appointment.ID,
			Time:   appointment.Time,
			Status: appointment.Status,
		})
	}
	return appointmentDTO, nil
}

func (a *AppointmentUsecase) Create(payload domain.AppointmentRequest, user *models.User) (*models.Appointment, error) {
	doctorUsecase := NewDoctorUsecase(*repository.NewDoctorRepository(a.AppointmentRepository.DB))
	if _, err := doctorUsecase.FindByID(payload.Doctor); err != nil {
		return nil, err
	}
	t, _ := time.Parse("2006-01-02 15:04", payload.Time)
	appointments, _ := a.AppointmentRepository.DoctorAppointmentBetween(payload.Doctor, t.Add(-10*time.Minute), t.Add(10*time.Minute))
	if len(appointments) > 0 {
		return nil, errors.New("time already appointment")
	}
	appointment, err := a.AppointmentRepository.Create(user.ID, payload.Doctor, t, "scheduled")
	if err != nil {
		return nil, err
	}
	return appointment, nil
}
