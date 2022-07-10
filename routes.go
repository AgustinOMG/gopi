package main

import "github.com/gin-gonic/gin"

func apiRoutes(router *gin.Engine) {

	router.GET("/kayser", getTodos)
	router.GET("/kayser/:id", getTodo)
	router.POST("/kayser", addTodo)
	router.PATCH("/kayser/:id", toggleTodoStatus)
	router.Run("localhost:9090")

}
