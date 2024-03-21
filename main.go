package main

import (
	"renderin-html/initializers"
	"renderin-html/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	routes.Routes()
	routes.ActionRoute()
}
