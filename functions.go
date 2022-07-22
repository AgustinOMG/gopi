package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func addPart(data *gin.Context) {

	stationData := make(map[string]string)
	id := data.Param("id")
	station := data.Param("station")
	line := data.Param("line")
	if err := data.BindJSON(&stationData); err != nil {
		data.IndentedJSON(http.StatusNotImplemented, gin.H{"mensaje": "Formato de datos incorrecto"})
	}
	dbStation := Client.Database("kayser").Collection(line)
	dbStation.InsertOne(context.TODO(), bson.D{{Key: "sn", Value: id}})
	stationData["station"] = station
	result, err := dbStation.UpdateOne(context.TODO(), bson.D{{Key: "sn", Value: id}, {Key: "station.station", Value: station}}, bson.D{{Key: "$set", Value: bson.D{{Key: "station", Value: stationData}}}})
	if err != nil {
		data.IndentedJSON(http.StatusNotImplemented, gin.H{"mensaje": "Ocurrio un error al momento de registrar la pieza"})
	} else {
		data.IndentedJSON(http.StatusOK, result)
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
