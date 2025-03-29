package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
	"github.com/JscorpTech/jst-go/api/middlewares"
	"github.com/JscorpTech/jst-go/repository"
)

func (r *Route) InitAppointmentRoute() {
	controller := controllers.NewAppointmentController(r.App)
	router := r.App.Server.Group("/appointment")
	router.GET("/", controller.List)
	router.POST("/", controller.Create, middlewares.AuthMiddleware(repository.NewUserRepository(r.App.DB)))
}
