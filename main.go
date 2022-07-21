package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://192.168.1.116:27017"

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

var Client *mongo.Client
var err error

func main() {
	// Crear la conexion con el servidor.
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Print(err)
	}
	defer func() {
		if err = Client.Disconnect(context.TODO()); err != nil {
			fmt.Print(err)
		}
	}()
	// Ping the primary
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Print(err)
	}

	usersCollection := Client.Database("testing").Collection("users")
	// insert a single document into a collection
	// create a bson.D object
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		fmt.Print(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
	// crear el objeto de ruteo para la REST API
	router := gin.Default()
	// LLamada a funcion de paso de rutas, el objeto debe ser pasado como un parametro
	apiRoutes(router)
	// Corre el servidor, especificar IP y Puerto deseado.
	router.Run("localhost:9090")

}
