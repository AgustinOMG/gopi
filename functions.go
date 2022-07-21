package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func addPart(data *gin.Context) {
	station := data.Param("station")
	id := data.Param("id")
	dbStation := Client.Database("kayser").Collection(station)
	var part bson.M
	err := dbStation.FindOne(context.TODO(), bson.M{"id": id}).Decode(&part)
	if err != nil {
		result, err := dbStation.InsertOne(context.TODO(), bson.M{"id": id, "fecha": time.Now(), "estado": "OK"})
		if err != nil {
			data.IndentedJSON(http.StatusNotImplemented, gin.H{"mensaje": "Ocurrio un error al momento de registrar la pieza"})
		} else {
			data.IndentedJSON(http.StatusOK, result)
		}
	} else {
		data.IndentedJSON(http.StatusAlreadyReported, gin.H{"mensaje": "La Pieza ya ha sido registrada en esta estacion."})

	}

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
