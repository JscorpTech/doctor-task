package controllers

import (
	"net/http"

	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	App *bootstrap.App
}

func NewDoctorController(app *bootstrap.App) *DoctorController {
	return &DoctorController{
		App: app,
	}
}

func (d *DoctorController) List(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"detail": "OK"})
}
