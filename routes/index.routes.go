package routes

import (
	"net/http"
	deliverytask "renderin-html/routes/delivery_task"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	router := app
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Connected to server",
		})
	})

	deliverytask.ActionRoute(app)

}
