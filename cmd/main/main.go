package main

import (
	"tudo-tools/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(HtmxMiddleware())

	router.HTMLRender = &TemplRender{}
	router.Static("/views/assets", "./views/assets")

	homeHandler := handlers.HomeHandler{}
	convertHandler := handlers.ConvertHandler{}
	metrosHandler := handlers.MetrosHandler{}

	router.GET("/", homeHandler.GetHome)
	router.GET("/convert/distance", convertHandler.GetDistance)
	router.GET("/convert/time", convertHandler.GetTime)
	router.POST("/convert", convertHandler.Convert)
	router.GET("/Metros", metrosHandler.GetMetros)

	err := router.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
