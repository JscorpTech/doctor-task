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

type AuthController struct {
	App        *bootstrap.App
	LoginUC    domain.LoginUsecase
	RegisterUC domain.RegisterUsecase
}

func NewAuthController(app *bootstrap.App) *AuthController {
	userRepository := repository.NewUserRepository(app.DB)
	return &AuthController{
		App:        app,
		LoginUC:    usecase.NewLoginUsecase(userRepository),
		RegisterUC: usecase.NewRegisterUsecase(userRepository),
	}
}

func (auth *AuthController) Login(c echo.Context) error {
	var payload domain.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.LoginUC.Login(payload.Phone, payload.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, domain.ErrorResponse(messages.PermissionDenied, nil))
	}
	token, err := auth.LoginUC.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, domain.ErrorResponse(messages.PermissionDenied, nil))
	}
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", token))
}

func (auth *AuthController) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implement!!!")
}

func (auth *AuthController) Register(c echo.Context) error {
	var payload domain.RegisterRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.BadRequest, nil))
	}

	if err := validator.ValidateRequest(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.ValidationError, err))
	}

	user, err := auth.RegisterUC.CreateUserIfNotExist(payload.Phone, payload.FirstName, payload.LastName, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse(messages.InternalError, domain.ValidationError{
			Field:   "phone",
			Type:    "unique",
			Message: messages.UserAlreadyExist,
		}))
	}

	token, err := auth.LoginUC.GetToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse(messages.InternalError, err))
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", domain.RegisterResponse{
		User: domain.User{
			ID:        user.ID,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: *token,
	}))
}

func (auth *AuthController) User(c echo.Context) error {
	user, _ := c.Get("user").(*models.User)
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", &domain.User{
		ID:        user.ID,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}))
}

func (auth *AuthController) ResendCode(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", nil))
}

func (auth *AuthController) ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.SuccessResponse("OK", nil))
}
