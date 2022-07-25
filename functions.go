package main

import (
	"context"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func addPart(data *gin.Context) {
	var stationList = [5]string{"110", "120", "130", "140", "150"}

	//declara variable para alamcenar datos en formato Map
	stationData := make(map[string]string)
	//Obtener el valor del ID de acuerdo al parametro de la URL
	id := data.Param("id")
	//Obtener ek parametro de la estacion
	station := data.Param("station")
	//Obtener el parametro de la linea de produccion
	line := data.Param("line")
	//Obtener los datos del ormato json y pasarlos a la variable de tipo mapa, si existe un error enviar mensaje de regreso
	if err := data.BindJSON(&stationData); err != nil {
		data.IndentedJSON(http.StatusNotImplemented, gin.H{"mensaje": "Formato de datos incorrecto"})
	}
	// Declaracion de la coleccion y base de datos a donde se quiere escribir
	dbStation := Client.Database("kayser").Collection(line)

	//Busca si existe ya una valor con ese Id.
	errid := dbStation.FindOne(context.TODO(), bson.M{"sn": id}).Err()

	if errid != nil {
		dbStation.InsertOne(context.TODO(), bson.M{"sn": id, "index": 0})
	}

	// agrega el valor de la estacion como una lista
	stationData["station"] = station

	// checa que la estacion anteriori haya sido registrada
	if station == stationList[0] {
		dbStation.UpdateOne(context.TODO(), bson.M{"sn": id}, bson.D{{Key: "$set", Value: bson.D{{Key: station, Value: stationData}, {Key: "index", Value: 1}}}})
		data.IndentedJSON(http.StatusNotImplemented, gin.H{"error": "nil", "mensaje": "Pieza Registrada con exito", "data": stationData})
	} else {
		var stationIndex bson.M
		dbStation.FindOne(context.TODO(), bson.M{"sn": id}).Decode(&stationIndex)
		index := stationIndex["index"].(int32)
		var indexInc int32
		if index <= 4 {
			indexInc = index + 1
		} else {
			indexInc = index
		}
		fmt.Print(index)
		if station == stationList[index] {
			dbStation.UpdateOne(context.TODO(), bson.M{"sn": id}, bson.D{{Key: "$set", Value: bson.D{{Key: station, Value: stationData}, {Key: "index", Value: indexInc}}}})
			data.IndentedJSON(http.StatusNotImplemented, gin.H{"error": "nil", "mensaje": "Pieza Registrada con exito", "data": stationData})
		} else {
			data.IndentedJSON(http.StatusNotImplemented, gin.H{"error": "Ksr-Stn-01", "mensaje": "La pieza no ha sido registrada en las estaciones necesarias"})
		}

	}

}

func getData(data *gin.Context) {
	line := data.Param("line")
	id := data.Param("id")
	dbStation := Client.Database("kayser").Collection(line)
	var lineData bson.M
	dbStation.FindOne(context.TODO(), bson.M{"sn": id}).Decode(&lineData)
	data.IndentedJSON(http.StatusCreated, lineData)
}
