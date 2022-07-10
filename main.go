package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Item: "first", Completed: false},
	{Id: "2", Item: "second", Completed: false},
	{Id: "3", Item: "third", Completed: false},
}

func main() {
	router := gin.Default()
	apiRoutes(router)
	router.Run("localhost:9090")

}
