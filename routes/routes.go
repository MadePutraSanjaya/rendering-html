package routes

import (
	"net/http"
	"renderin-html/initializers"
	"renderin-html/models"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {

		var deliverys []models.Delivery
		if err := initializers.DB.Find(&deliverys).Error; err != nil {
            c.HTML(http.StatusInternalServerError, "error.html", gin.H{
                "error": err.Error(),
            })
            return
        }

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"delivery" : deliverys,
		})
	})

	r.Run()
}
