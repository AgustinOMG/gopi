package main

import "github.com/gin-gonic/gin"

func apiRoutes(router *gin.Engine) {

	router.POST("/ksr/:line/:station/:id", addPart)

	router.GET("/ksr/getdata/:line/:id", getData)

}
