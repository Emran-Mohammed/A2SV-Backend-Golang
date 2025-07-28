package models

import (
	domain "task_manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskMongoModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DueDate     time.Time          `bson:"due_date"`
	Status      string             `bson:"status"`
}


func DomainToMongoTask(d *domain.Task)(*TaskMongoModel,error){
	var id primitive.ObjectID
    var err error
	if d.ID != ""{
		id, err = primitive.ObjectIDFromHex(d.ID)
		if err != nil{
			return nil, err
		}
	}


return &TaskMongoModel{
	ID: id,
	Title: d.Title,
	Description: d.Description,
	DueDate:     d.DueDate,
	Status:      string(d.Status),
}, nil
}

func MongoToDomainTask(m *TaskMongoModel) *domain.Task {
	return &domain.Task{
		ID:          m.ID.Hex(),
		Title:       m.Title,
		Description: m.Description,
		DueDate:     m.DueDate,
		Status:      domain.Status(m.Status),
	}
}
