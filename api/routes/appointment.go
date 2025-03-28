package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
)

func (r *Route) InitAppointmentRoute() {
	controller := controllers.NewAppointmentController(r.App)
	router := r.App.Server.Group("/appointment")
	router.GET("/", controller.List)
}
