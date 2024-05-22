package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-api/handlers"
)

func TestGetToDoItems(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Couldn't connect to mongo: %v\n", err)
	}
	collection := client.Database("test").Collection("todo")

	h := &handlers.Handler{
		Collection: collection,
	}

	r := gin.Default()
	r.GET("/todo", h.GetToDoItems)

	req, err := http.NewRequest(http.MethodGet, "/todo", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status OK; got %v\n", resp.Code)
	}
}
