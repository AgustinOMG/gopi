package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodos(data *gin.Context) {
	data.IndentedJSON(http.StatusOK, todos)

}

func addTodo(data *gin.Context) {
	var newTodo todo
	if err := data.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	data.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo Not Found")
}

func getTodo(data *gin.Context) {
	id := data.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		data.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}

	data.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(data *gin.Context) {
	id := data.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		data.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}

	todo.Completed = !todo.Completed
	data.IndentedJSON(http.StatusOK, todo)

}
