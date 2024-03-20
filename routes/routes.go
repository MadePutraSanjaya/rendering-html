package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.Run()
}
