package deliverytask

import (
	"renderin-html/controllers"

	"github.com/gin-gonic/gin"
)

func ActionRoute(app *gin.Engine) {
	r := app
	r.GET("/deliveries", controllers.DeliverysIndex)
	r.GET("/deliveries/:id", controllers.DeliverysShow)
	r.POST("/deliveries", controllers.DeliverysCreate)
	r.PUT("/deliveries/:id", controllers.DeliverysUpdate)
	r.DELETE("/deliveries/:id", controllers.DeliverysDelete)

}
