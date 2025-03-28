package controllers

import (
	"net/http"

	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/labstack/echo/v4"
)

type AppointmentController struct {
	App *bootstrap.App
}

func NewAppointmentController(app *bootstrap.App) *AppointmentController {
	return &AppointmentController{
		App: app,
	}
}

func (d *AppointmentController) List(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"detail": "OK"})
}
