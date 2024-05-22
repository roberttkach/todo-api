package models

type ToDoItem struct {
	ID        string `bson:"_id"`
	Text      string `bson:"text"`
	Timestamp string `bson:"timestamp"`
	List      string `bson:"list"`
}
