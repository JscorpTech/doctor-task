package domain

import (
	"time"

	"github.com/JscorpTech/jst-go/models"
)

type AppointmentResponse struct {
	ID     uint      `json:"id"`
	Time   time.Time `json:"time"`
	Status string    `json:"status"` // "scheduled", "cancelled"
}

type AppointmentDetailResponse struct {
	AppointmentResponse
	Doctor DoctorResponse `json:"doctor"`
}

type AppointmentRequest struct {
	Doctor uint   `json:"doctor" validate:"required"`
	Time   string `json:"time" validate:"required,datetime=2006-01-02 15:04"`
}

type AppointmentUsecase interface {
	List() ([]models.Appointment, error)
	Create(AppointmentRequest, *models.User) (*models.Appointment, error)
}
