package routes

import (
	"github.com/JscorpTech/jst-go/api/controllers"
)

func (r *Route) InitDoctorRoute() {
	controller := controllers.NewDoctorController(r.App)
	router := r.App.Server.Group("/doctor")
	router.GET("/", controller.List)
	router.GET("/:id/", controller.Retrieve)
}
