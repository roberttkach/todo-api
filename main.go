package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-api/handlers"
)

func main() {
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)

	log.SetOutput(logFile)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("test").Collection("todo")

	h := &handlers.Handler{
		Collection: collection,
	}

	r := gin.Default()

	r.GET("/todo", h.GetToDoItems)
	r.POST("/todo", h.CreateToDoItem)
	r.PUT("/todo/:id", h.UpdateToDoItem)
	r.DELETE("/todo/:id", h.DeleteToDoItem)

	err = r.Run()
	if err != nil {
		return
	}
}
