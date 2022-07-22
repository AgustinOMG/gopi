package main

import "github.com/gin-gonic/gin"

func apiRoutes(router *gin.Engine) {

	router.POST("/ksr/:line/:station/:id", addPart)
	router.GET("/monbai/:id", getTodo)
	router.POST("/monbai", addTodo)
	router.PATCH("/monbai/:id", toggleTodoStatus)
	router.Run("localhost:9090")

}
