package controllers

import (
	"net/http"
	"strconv"

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
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", d.DoctorUsecase.List(c.QueryParam("search"))))
}

func (d *DoctorController) Retrieve(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.ParseUint(idString, 10, 32)
	doctor, err := d.DoctorUsecase.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrorResponse("Doctor Not found", nil))
	}
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", doctor))
}
