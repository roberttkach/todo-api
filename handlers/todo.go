package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-api/models"
)

type Handler struct {
	Collection *mongo.Collection
}

func (h *Handler) GetToDoItems(c *gin.Context) {
	list := c.Query("list")
	sort := c.DefaultQuery("sort", "timestamp")

	var results []*models.ToDoItem
	opts := options.Find().SetSort(bson.D{{Key: sort, Value: 1}})
	cur, err := h.Collection.Find(context.TODO(), bson.M{"list": list}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ToDo items"})
		return
	}
	defer func() {
		if err = cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	for cur.Next(context.TODO()) {
		var elem models.ToDoItem
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			continue
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ToDo items"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *Handler) CreateToDoItem(c *gin.Context) {
	var todo models.ToDoItem
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Collection.InsertOne(context.TODO(), todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating ToDo item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo item created", "id": res.InsertedID})
}

func (h *Handler) UpdateToDoItem(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDoItem
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"text":      todo.Text,
			"timestamp": todo.Timestamp,
			"list":      todo.List,
		},
	}

	res, err := h.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating ToDo item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo item updated", "matched_count": res.MatchedCount, "modified_count": res.ModifiedCount})
}

func (h *Handler) DeleteToDoItem(c *gin.Context) {
	id := c.Param("id")

	res, err := h.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting ToDo item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo item deleted", "deleted_count": res.DeletedCount})
}
