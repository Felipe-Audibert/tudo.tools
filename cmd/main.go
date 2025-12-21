package main

import (
	"tudo-tools/handers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.HTMLRender = &TemplRender{}

	router.Static("/views/assets", "./views/assets")

	homeHandler := handers.HomeHandler{}

	router.GET("/", homeHandler.GetHome)

	err := router.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
