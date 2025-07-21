package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Task struct{
	ID			primitive.ObjectID	`json:"id" bson:"_id,omitempity"`
	Title		string 		`json:"title" binding:"required" bson:"title"`
	Description string 		`json:"description" bson:"description"`
	DueDate		time.Time  	`json:"duedate" binding:"required" bson:"duedate"`
	Status 		Status		`json:"status" binding:"required" bson:"status"`

}

type Status string

const(
	Progress Status = "progress"
	Completed Status = "completed"
	Pending Status = "pending"
)
func (s Status) IsValid() bool {
    return s == Progress || s == Completed || s == Pending
}