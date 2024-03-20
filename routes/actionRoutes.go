package routes

import (
	"renderin-html/controllers"

	"github.com/gin-gonic/gin"
)

func ActionRoute() {
	r := gin.Default()

	r.GET("/delivery", controllers.DeliverysIndex)
	r.GET("/delivery/:id", controllers.DeliverysShow)
	r.DELETE("/delivery/:id", controllers.DeliverysDelete)
	r.PUT("/delivery/:id", controllers.DeliverysUpdate)
	r.POST("/delivery", controllers.DeliverysCreate)

	r.Run()

}