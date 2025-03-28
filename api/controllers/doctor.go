package controllers

import (
	"net/http"

	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/domain"
	"github.com/JscorpTech/jst-go/repository"
	"github.com/JscorpTech/jst-go/usecase"
	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	App           *bootstrap.App
	DoctorUsecase domain.DoctorUsecase
}

func NewDoctorController(app *bootstrap.App) *DoctorController {
	return &DoctorController{
		App:           app,
		DoctorUsecase: usecase.NewDoctorUsecase(*repository.NewDoctorRepository(app.DB)),
	}
}

func (d *DoctorController) List(c echo.Context) error {
	doctors, _ := d.DoctorUsecase.List()
	return c.JSON(http.StatusOK, doctors)
}
