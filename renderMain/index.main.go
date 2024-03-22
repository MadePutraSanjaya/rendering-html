package rendermain

import (
	"renderin-html/initializers"
	routes "renderin-html/routes"

	"github.com/gin-gonic/gin"
)

func MainRoot() {
	initializers.ConnectToDb()
	initializers.LoadEnvVariables()

	app := gin.Default()
	app.LoadHTMLGlob("views/**/*.html")
	app.Static("/static", "./static")
	routes.Routes(app)
	app.Run()
}
