package main

import "github.com/gin-gonic/gin"

func apiRoutes(router *gin.Engine) {

	router.GET("/monbai", getTodos)
	router.GET("/monbai/:id", getTodo)
	router.POST("/monbai", addTodo)
	router.PATCH("/monbai/:id", toggleTodoStatus)
	router.Run("localhost:9090")

}
