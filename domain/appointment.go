package domain

import "github.com/JscorpTech/jst-go/models"

type AppointmentResponse struct {
}

type AppointmentUsecase interface {
	List() ([]*models.Appointment, error)
}
