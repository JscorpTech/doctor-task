package controllers

import (
	"net/http"

	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/models"
	"github.com/JscorpTech/jst-go/pkg/messages"
	"github.com/JscorpTech/jst-go/pkg/validator"
	"github.com/JscorpTech/jst-go/repository"
	"github.com/JscorpTech/jst-go/usecase"
	"github.com/labstack/echo/v4"
)

type AppointmentController struct {
	App           *bootstrap.App
	AppointmentUC domain.AppointmentUsecase
}

func NewAppointmentController(app *bootstrap.App) *AppointmentController {
	appointmentRepository := repository.NewAppointmentRepository(app.DB)
	return &AppointmentController{
		App:           app,
		AppointmentUC: usecase.NewAppointmentUsecase(*appointmentRepository),
	}
}

func (a *AppointmentController) List(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"detail": "OK"})
}

func (a *AppointmentController) Create(c echo.Context) error {
	var payload domain.AppointmentRequest
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.BadRequest, err))
	}
	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}
	appointment, err := a.AppointmentUC.Create(payload, c.Get("user").(*models.User))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, domain.SuccessResponse("Appointment created", domain.AppointmentResponse{
		ID:     appointment.ID,
		Time:   appointment.Time,
		Status: appointment.Status,
	}))
}
